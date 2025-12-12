package app

import (
	"bufio"
	"os"
	"strings"
	"time"

	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

// DownloadOptions 下载选项
type DownloadOptions struct {
	Tasks              []string                        // RJ号列表
	ShowProgress       bool                            // 是否显示进度信息
	WaitForInput       bool                            // 完成后是否等待用户输入
	OnTimeout          func()                          // 超时回调
	TimeoutDetected    chan bool                       // 超时检测通道
	FormatFilter       *spider.FilterStrategy          // 格式过滤策略
	OnFormatConflict   func(*spider.FormatAnalysis, func(*spider.FilterStrategy)) *spider.FilterStrategy // 格式冲突回调（带保存回调）
}

// DownloadResult 下载结果
type DownloadResult struct {
	Success      bool
	FailedTasks  []spider.FailedTask
	TimedOut     bool
}

// Download 执行下载任务 (CLI和UI共用的核心逻辑)
func (a *App) Download(opts DownloadOptions) *DownloadResult {
	result := &DownloadResult{
		Success: true,
	}

	// 使用配置文件中的设置
	c := spider.NewASMRClient(
		a.Config.MaxTask,
		a.Config.MaxThread,
		a.Config.MaxRetry,
		a.Config.GetBufferSize(),
	)
	c.WorkerPool.Start()

	if opts.ShowProgress {
		utils.Info(i18n.T("using_threads", a.Config.MaxThread, a.Config.MaxTask, a.Config.MaxRetry))
	}

	// 登录
	err := c.Login()
	if err != nil {
		utils.Error(i18n.T("login_failed", err))
		if opts.WaitForInput {
			utils.Warning(i18n.T("press_enter_to_return"))
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')
		}
		result.Success = false
		result.TimedOut = false
		return result
	}

	// 转换回调函数类型（将带保存回调的格式转换为spider需要的简单格式）
	var convertedCallback func(*spider.FormatAnalysis) *spider.FilterStrategy
	if opts.OnFormatConflict != nil {
		convertedCallback = func(analysis *spider.FormatAnalysis) *spider.FilterStrategy {
			// Download函数不负责保存策略，传递空回调
			return opts.OnFormatConflict(analysis, func(*spider.FilterStrategy) {})
		}
	}

	// 提交下载任务
	for _, task := range opts.Tasks {
		c.DownloadWithFilter(task, opts.FormatFilter, convertedCallback)
	}

	// 等待任务完成
	done := make(chan bool)
	go func() {
		c.WorkerPool.Wait()

		// 重试失败的任务
		for c.RetryFailedTasks() {
			c.WorkerPool.Wait()
		}
		done <- true
	}()

	// 等待完成或超时
	if opts.TimeoutDetected != nil {
		select {
		case <-done:
			// 正常完成
			result.FailedTasks = c.FailedTasks
			result.Success = true
			result.TimedOut = false
		case <-opts.TimeoutDetected:
			// 检测到超时
			result.Success = false
			result.TimedOut = true
		}
	} else {
		// 无超时检测,直接等待完成
		<-done
		result.FailedTasks = c.FailedTasks
		result.Success = true
		result.TimedOut = false
	}

	// 显示结果
	if result.Success && opts.ShowProgress {
		if len(result.FailedTasks) > 0 {
			utils.Error( i18n.T("download_failed_count", len(result.FailedTasks)))
			utils.Info(i18n.T("failed_files_list"))
			for _, task := range result.FailedTasks {
				utils.Errorf("  - %s", task.FileName)
			}
		} else {
			utils.Success(i18n.T("download_complete"))
		}
	}

	return result
}

// 带超时监控和格式选择的下载
func (a *App) DownloadWithMonitorAndFormatSelector(tasks []string, onFormatConflict func(*spider.FormatAnalysis, func(*spider.FilterStrategy)) *spider.FilterStrategy) {
	// 先暂时保存下载状态（不含格式策略）
	if err := config.SaveDownloadState(a.Config, tasks); err != nil {
		utils.Error(i18n.T("save_download_state_failed", err))
	}

	// 启动活动监控
	utils.GlobalMonitor.Start()
	defer utils.GlobalMonitor.Stop()

	// 创建保存回调函数，供格式选择器在每次用户选择后立即保存
	saveStrategyCallback := func(strategy *spider.FilterStrategy) {
		if strategy != nil {
			configStrategy := convertToConfigStrategy(strategy)
			if err := config.SaveDownloadStateWithStrategy(a.Config, tasks, configStrategy); err != nil {
				utils.Error(i18n.T("save_download_state_failed", err))
			}
		}
	}

	// 创建一个包装的回调函数，注入保存回调
	wrappedCallback := func(analysis *spider.FormatAnalysis, _ func(*spider.FilterStrategy)) *spider.FilterStrategy {
		if onFormatConflict != nil {
			// 调用格式选择器，传入保存回调
			return onFormatConflict(analysis, saveStrategyCallback)
		}
		return nil
	}

	// 使用循环而不是递归，避免栈溢出
	for {
		// 启动超时检测goroutine
		stopMonitor := make(chan bool)
		timeoutDetected := make(chan bool, 1)

		go func() {
			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if utils.GlobalMonitor.IsTimeout() {
						inactiveTime := utils.GlobalMonitor.GetInactiveTime()
						utils.Warning(i18n.T("timeout_detected", inactiveTime.Round(time.Second)))
						timeoutDetected <- true
						return
					}
				case <-stopMonitor:
					return
				}
			}
		}()

		// 执行下载
		result := a.Download(DownloadOptions{
			Tasks:            tasks,
			ShowProgress:     true,
			WaitForInput:     false,
			TimeoutDetected:  timeoutDetected,
			OnFormatConflict: wrappedCallback,
		})

		// 停止监控
		close(stopMonitor)

		// 如果检测到超时，重启下载（循环继续）
		if result.TimedOut {
			utils.Warning(i18n.T("download_interrupted"))
			time.Sleep(2 * time.Second)
			continue // 继续循环重试
		}

		// 下载成功，跳出循环
		break
	}

	// 清除下载状态
	if err := config.ClearDownloadState(a.Config); err != nil {
		utils.Error(i18n.T("clear_download_state_failed", err))
	}

	// 等待用户输入
	utils.Info(i18n.T("press_enter_to_return"))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

// 检查并处理未完成的下载
func (a *App) CheckUnfinishedDownload(onContinue func([]string, *config.FormatStrategy)) {
	if a.Config.DownloadState.InProgress && len(a.Config.DownloadState.Tasks) > 0 {
		utils.Warning(i18n.T("unfinished_download_detected"))
		utils.Info(i18n.T("unfinished_tasks", strings.Join(a.Config.DownloadState.Tasks, ", ")))

		// 显示已保存的格式策略信息
		if a.Config.DownloadState.FormatStrategy != nil {
			utils.Info(i18n.T("format_strategy_saved"))
		}

		reader := bufio.NewReader(os.Stdin)
		utils.Info(i18n.T("continue_download_prompt"))
		input, _ := reader.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(input))

		if choice == "y" || choice == "yes" {
			utils.Info(i18n.T("continuing_download"))
			if onContinue != nil {
				onContinue(a.Config.DownloadState.Tasks, a.Config.DownloadState.FormatStrategy)
			}
		} else {
			utils.Info(i18n.T("download_cancelled"))
			config.ClearDownloadState(a.Config)
		}
	}
}

func convertToConfigStrategy(spiderStrategy *spider.FilterStrategy) *config.FormatStrategy {
	if spiderStrategy == nil {
		return nil
	}
	return &config.FormatStrategy{
		Mode:             spiderStrategy.Mode,
		PriorityFormats:  spiderStrategy.PriorityFormats,
		ManualSelections: spiderStrategy.ManualSelections,
		IncludeFormats:   spiderStrategy.IncludeFormats,
	}
}
