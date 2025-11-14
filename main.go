package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	utils.Success(i18n.T("welcome", i18n.AppName()))
	utils.Info(i18n.T("config_loaded"))

	// 检查是否有未完成的下载
	checkUnfinishedDownload()

	for {
		showMenu()
		choice := readInput(i18n.T("prompt_choose_operation"))
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			startDownload()
		case "2":
			modifyConfig()
		case "3":
			showAbout()
		case "4":
			changeLanguage()
		case "0":
			utils.Info(i18n.T("goodbye"))
			return
		default:
			utils.Warning(i18n.T("invalid_choice"))
		}
	}
}

func checkUnfinishedDownload() {
	if spider.Conf.DownloadState.InProgress && len(spider.Conf.DownloadState.Tasks) > 0 {
		utils.Warning(i18n.T("unfinished_download_detected"))
		utils.Info(i18n.T("unfinished_tasks", strings.Join(spider.Conf.DownloadState.Tasks, ", ")))

		choice := readInput(i18n.T("continue_download_prompt"))
		choice = strings.ToLower(strings.TrimSpace(choice))

		if choice == "y" || choice == "yes" {
			utils.Info(i18n.T("continuing_download"))
			continueDownload(spider.Conf.DownloadState.Tasks)
		} else {
			utils.Info(i18n.T("download_cancelled"))
			// 清除下载状态
			config.ClearDownloadState(spider.Conf)
		}
	}
}

func showMenu() {
	utils.Info(strings.Repeat("=", 50))
	utils.Info(i18n.T("main_menu"))
	utils.Info(strings.Repeat("=", 50))
	utils.Info("  " + i18n.T("menu_start_download"))
	utils.Info("  " + i18n.T("menu_modify_config"))
	utils.Info("  " + i18n.T("menu_about"))
	utils.Info("  " + i18n.T("menu_language") + " / Language")
	utils.Info("  " + i18n.T("menu_exit"))
	utils.Info(strings.Repeat("=", 50))
}

func readInput(prompt string) string {
	fmt.Printf("%s: ", prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func startDownload() {
	utils.Info(i18n.T("start_download_title"))
	utils.Info(i18n.T("download_input_hint"))
	utils.Info(i18n.T("download_multiple_hint"))

	rjNumbers := readInput(i18n.T("prompt_rj_number"))
	if rjNumbers == "" {
		utils.Warning(i18n.T("no_rj_input"))
		return
	}

	tasks := strings.Split(rjNumbers, " ")
	executeDownload(tasks)
}

func continueDownload(tasks []string) {
	executeDownload(tasks)
}

func executeDownload(tasks []string) {
	// 保存下载状态
	if err := config.SaveDownloadState(spider.Conf, tasks); err != nil {
		utils.Error(i18n.T("save_download_state_failed", err))
	}

	// 启动活动监控
	utils.GlobalMonitor.Start()
	defer utils.GlobalMonitor.Stop()

	// 启动超时检测goroutine
	stopMonitor := make(chan bool)
	timeoutDetected := make(chan bool, 1)

	go func() {
		ticker := time.NewTicker(30 * time.Second) // 每30秒检查一次
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
	success := performDownload(tasks, timeoutDetected)

	// 停止监控
	close(stopMonitor)

	// 如果检测到超时，重启下载
	if !success {
		utils.Warning(i18n.T("download_interrupted"))
		time.Sleep(2 * time.Second)
		executeDownload(tasks) // 递归重启
		return
	}

	// 清除下载状态
	if err := config.ClearDownloadState(spider.Conf); err != nil {
		utils.Error(i18n.T("clear_download_state_failed", err))
	}

	utils.Info(i18n.T("press_enter_to_return"))
	reader.ReadString('\n')
}

func performDownload(tasks []string, timeoutChan <-chan bool) bool {
	// 使用配置文件中的设置
	c := spider.NewASMRClient(spider.Conf.MaxTask, spider.Conf.MaxThread, spider.Conf.MaxRetry)
	c.WorkerPool.Start()

	utils.Info(i18n.T("using_threads", spider.Conf.MaxThread, spider.Conf.MaxTask, spider.Conf.MaxRetry))

	err := c.Login()
	if err != nil {
		utils.Error(i18n.T("login_failed", err))
		utils.Warning(i18n.T("press_enter_to_return"))
		reader.ReadString('\n')
		return true // 登录失败不算超时，返回true结束
	}

	for _, task := range tasks {
		c.Download(task)
	}

	// 等待任务完成，同时监听超时信号
	done := make(chan bool)
	go func() {
		c.WorkerPool.Wait()

		// 重试失败的任务，直到全部成功或达到最大重试次数
		for c.RetryFailedTasks() {
			c.WorkerPool.Wait()
		}
		done <- true
	}()

	// 等待完成或超时
	select {
	case <-done:
		// 正常完成
		if len(c.FailedTasks) > 0 {
			utils.Error(i18n.T("download_failed_count", len(c.FailedTasks)))
			utils.Info(i18n.T("failed_files_list"))
			for _, task := range c.FailedTasks {
				utils.Error("  - %s", task.FileName)
			}
		} else {
			utils.Success(i18n.T("download_complete"))
		}
		return true
	case <-timeoutChan:
		// 检测到超时
		return false
	}
}

func modifyConfig() {
	utils.Info(i18n.T("modify_config_title"))

	cfg := spider.Conf

	utils.Info("\n" + i18n.T("current_config"))
	utils.Info("  " + i18n.T("config_account", cfg.Account))
	utils.Info("  " + i18n.T("config_password", strings.Repeat("*", len(cfg.Password))))
	utils.Info("  " + i18n.T("config_max_task", cfg.MaxTask))
	utils.Info("  " + i18n.T("config_max_thread", cfg.MaxThread))
	utils.Info("  " + i18n.T("config_max_retry", cfg.MaxRetry))
	utils.Info("  " + i18n.T("config_language", i18n.GetLocaleName(cfg.Language)))
	proxyDisplay := cfg.Proxy
	if proxyDisplay == "" {
		proxyDisplay = i18n.T("proxy_disabled")
	}
	utils.Info("  " + i18n.T("config_proxy", proxyDisplay))
	utils.Info("  " + i18n.T("config_return"))

	choice := readInput("\n" + i18n.T("prompt_choose_modify"))

	switch choice {
	case "1":
		newAccount := readInput(i18n.T("prompt_new_account"))
		if newAccount != "" {
			cfg.Account = newAccount
			utils.Success(i18n.T("account_updated"))
		}
	case "2":
		newPassword := readInput(i18n.T("prompt_new_password"))
		if newPassword != "" {
			cfg.Password = newPassword
			utils.Success(i18n.T("password_updated"))
		}
	case "3":
		newMaxTask := readInput(i18n.T("prompt_new_max_task"))
		if val, err := strconv.Atoi(newMaxTask); err == nil && val > 0 {
			cfg.MaxTask = val
			utils.Success(i18n.T("max_task_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "4":
		newMaxThread := readInput(i18n.T("prompt_new_max_thread"))
		if val, err := strconv.Atoi(newMaxThread); err == nil && val > 0 {
			cfg.MaxThread = val
			utils.Success(i18n.T("max_thread_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "5":
		newMaxRetry := readInput(i18n.T("prompt_new_max_retry"))
		if val, err := strconv.Atoi(newMaxRetry); err == nil && val >= 0 {
			cfg.MaxRetry = val
			utils.Success(i18n.T("max_retry_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "6":
		changeLanguage()
		return
	case "7":
		newProxy := readInput(i18n.T("prompt_new_proxy"))
		newProxy = strings.TrimSpace(newProxy)
		if newProxy == "" {
			cfg.Proxy = ""
			utils.SetProxy("")
			utils.Success(i18n.T("proxy_disabled"))
		} else {
			if err := utils.SetProxy(newProxy); err != nil {
				utils.Error(i18n.T("invalid_proxy", err))
				return
			}
			cfg.Proxy = newProxy
			utils.Success(i18n.T("proxy_updated", newProxy))
		}
	case "0":
		return
	default:
		utils.Warning(i18n.T("invalid_choice"))
		return
	}

	// 保存配置
	if err := config.SaveConfig(cfg); err != nil {
		utils.Error(i18n.T("save_config_failed", err))
	} else {
		utils.Success(i18n.T("config_saved"))
	}
}

func changeLanguage() {
	utils.Info(i18n.T("language_title"))

	// 动态显示所有支持的语言
	languages := i18n.GetSupportedLocales()
	langMap := make(map[string]string)

	for idx, lang := range languages {
		num := idx + 1
		utils.Info(fmt.Sprintf("  %d. %s", num, lang.NativeName))
		langMap[strconv.Itoa(num)] = lang.Code
	}
	utils.Info("  " + i18n.T("language_return"))

	choice := readInput(i18n.T("prompt_choose_language"))
	choice = strings.TrimSpace(choice)

	if choice == "0" {
		return
	}

	newLang, ok := langMap[choice]
	if !ok {
		utils.Warning(i18n.T("invalid_choice"))
		return
	}

	// 更新配置
	spider.Conf.Language = newLang
	i18n.SetLocale(newLang)

	// 保存配置
	if err := config.SaveConfig(spider.Conf); err != nil {
		utils.Error(i18n.T("save_config_failed", err))
	} else {
		utils.Success(i18n.T("language_updated", i18n.GetLocaleName(newLang)))
	}
}

func showAbout() {
	utils.Info(i18n.T("about_title"))

	// 输出关于内容（所有格式都在翻译文件的list中定义）
	appName := i18n.AppName()
	version := i18n.Version()
	for _, line := range i18n.TList("about_content") {
		// 替换占位符为实际的值
		line = strings.ReplaceAll(line, "{{app_name}}", appName)
		line = strings.ReplaceAll(line, "{{version}}", version)
		utils.Info(line)
	}

	utils.Info(i18n.T("press_enter_to_return"))
	reader.ReadString('\n')
}
