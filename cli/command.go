package cli

import (
	"os"
	"strings"

	"re-asmr-spider/app"
	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

// Command CLI命令执行器
type Command struct {
	Flags *Flags
	App   *app.App
}

// NewCommand 创建命令执行器
func NewCommand(flags *Flags, application *app.App) *Command {
	return &Command{
		Flags: flags,
		App:   application,
	}
}

// ApplyFlags 应用命令行参数到配置
func (cmd *Command) ApplyFlags() {
	// 处理自定义配置文件路径
	if cmd.Flags.ConfigPath != "" {
		if _, err := os.Stat(cmd.Flags.ConfigPath); os.IsNotExist(err) {
			utils.Errorf("配置文件不存在: %s", cmd.Flags.ConfigPath)
			os.Exit(1)
		}
		utils.Infof("使用配置文件: %s", cmd.Flags.ConfigPath)

		// 从指定路径加载配置
		customCfg, err := config.GetConfigFromPath(cmd.Flags.ConfigPath)
		if err != nil {
			utils.Errorf("加载配置文件失败: %v", err)
			os.Exit(1)
		}
		// 更新应用配置和全局配置
		cmd.App.Config = customCfg
		spider.Conf = customCfg

		// 重新初始化i18n和代理（如果配置中有设置）
		if err := i18n.Init(customCfg.Language); err != nil {
			utils.Warningf("初始化语言设置失败: %v", err)
		}
		if customCfg.Proxy != "" {
			if err := utils.SetProxy(customCfg.Proxy); err != nil {
				utils.Warningf("设置代理失败: %v", err)
			}
		}
	}

	cfg := cmd.App.GetConfig()

	// 命令行参数覆盖配置文件
	if cmd.Flags.Account != "" {
		cfg.Account = cmd.Flags.Account
	}
	if cmd.Flags.Password != "" {
		cfg.Password = cmd.Flags.Password
	}
	if cmd.Flags.MaxTask > 0 {
		cfg.MaxTask = cmd.Flags.MaxTask
	}
	if cmd.Flags.MaxThread > 0 {
		cfg.MaxThread = cmd.Flags.MaxThread
	}
	if cmd.Flags.MaxRetry >= 0 {
		cfg.MaxRetry = cmd.Flags.MaxRetry
	}
	if cmd.Flags.BufferSizeMB > 0 {
		if cmd.Flags.BufferSizeMB >= 1 && cmd.Flags.BufferSizeMB <= 64 {
			cfg.BufferSizeMB = cmd.Flags.BufferSizeMB
		} else {
			utils.Warning("缓冲区大小应在 1-64 MB 之间，使用配置文件中的值")
		}
	}
	if cmd.Flags.Proxy != "" {
		if err := utils.SetProxy(cmd.Flags.Proxy); err != nil {
			utils.Errorf("无效的代理地址: %v", err)
			os.Exit(1)
		}
		cfg.Proxy = cmd.Flags.Proxy
	}
}

// Execute 执行CLI命令
func (cmd *Command) Execute() {
	utils.Success(i18n.T("welcome", i18n.AppName()))

	// 解析RJ号
	rjList := strings.Split(cmd.Flags.DownloadRJs, ",")
	tasks := make([]string, 0)
	for _, rj := range rjList {
		rj = strings.TrimSpace(rj)
		if rj != "" {
			tasks = append(tasks, rj)
		}
	}

	if len(tasks) == 0 {
		utils.Error("未指定有效的RJ号")
		os.Exit(1)
	}

	utils.Infof("命令行模式: 准备下载 %d 个任务", len(tasks))
	utils.Infof("RJ号: %s", strings.Join(tasks, ", "))

	// 构建格式过滤策略
	var filterStrategy *spider.FilterStrategy
	if cmd.Flags.FormatPriority != "" || cmd.Flags.IncludeFormats != "" {
		filterStrategy = &spider.FilterStrategy{
			Mode:            "priority",
			PriorityFormats: nil,
			IncludeFormats:  nil,
		}

		// 优先级模式
		if cmd.Flags.FormatPriority != "" {
			priorityFormats := cmd.parseFormatList(cmd.Flags.FormatPriority)
			filterStrategy.PriorityFormats = priorityFormats
			utils.Info(i18n.T("format_priority_applied", strings.Join(priorityFormats, " > ")))
		}

		// 额外包含的扩展名
		if cmd.Flags.IncludeFormats != "" {
			includeFormats := cmd.parseFormatList(cmd.Flags.IncludeFormats)
			filterStrategy.IncludeFormats = includeFormats
			utils.Infof("将额外下载以下扩展名的文件: %s", strings.Join(includeFormats, ", "))
		}
	}

	// 执行下载 (使用app包的共用逻辑)
	result := cmd.App.Download(app.DownloadOptions{
		Tasks:           tasks,
		ShowProgress:    true,
		WaitForInput:    false,
		TimeoutDetected: nil, // CLI模式不使用超时监控
		FormatFilter:    filterStrategy,
	})

	if result.Success {
		if len(result.FailedTasks) == 0 {
			utils.Success("所有下载任务完成")
			os.Exit(0)
		} else {
			utils.Warningf("部分任务失败: %d/%d", len(result.FailedTasks), len(tasks))
			os.Exit(1)
		}
	} else {
		utils.Error("下载过程中出现错误")
		os.Exit(1)
	}
}

// parseFormatList 解析格式列表字符串
func (cmd *Command) parseFormatList(formatStr string) []string {
	parts := strings.Split(formatStr, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		ext := strings.ToLower(strings.TrimSpace(part))
		if ext == "" {
			continue
		}
		// 确保有点号前缀
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		result = append(result, ext)
	}

	return result
}
