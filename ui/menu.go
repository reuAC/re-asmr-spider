package ui

import (
	"bufio"
	"os"
	"strings"

	"re-asmr-spider/app"
	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

// RunInteractive 运行交互模式
func RunInteractive(application *app.App) {
	utils.Success(i18n.T("welcome", i18n.AppName()))
	utils.Info(i18n.T("config_loaded"))

	// 检查未完成的下载
	application.CheckUnfinishedDownload(func(tasks []string, savedStrategy *config.FormatStrategy) {
		// 如果有保存的格式策略，询问用户是否使用
		var formatCallback func(*spider.FormatAnalysis, func(*spider.FilterStrategy)) *spider.FilterStrategy

		if savedStrategy != nil {
			// 提示用户选择
			utils.Info(i18n.T("format_strategy_found"))
			utils.Info(i18n.T("use_saved_strategy_prompt"))
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			choice := strings.ToLower(strings.TrimSpace(input))

			if choice == "y" || choice == "yes" || choice == "" {
				// 使用保存的策略
				utils.Success(i18n.T("using_saved_strategy"))
				spiderStrategy := &spider.FilterStrategy{
					Mode:             savedStrategy.Mode,
					PriorityFormats:  savedStrategy.PriorityFormats,
					ManualSelections: savedStrategy.ManualSelections,
				}
				formatCallback = func(analysis *spider.FormatAnalysis, saveCallback func(*spider.FilterStrategy)) *spider.FilterStrategy {
					return spiderStrategy
				}
			} else {
				// 重新选择
				utils.Info(i18n.T("reselecting_format"))
				formatCallback = func(analysis *spider.FormatAnalysis, saveCallback func(*spider.FilterStrategy)) *spider.FilterStrategy {
					return BuildFormatSelectorCallback(saveCallback)(analysis)
				}
			}
		} else {
			// 没有保存的策略，正常选择
			formatCallback = func(analysis *spider.FormatAnalysis, saveCallback func(*spider.FilterStrategy)) *spider.FilterStrategy {
				return BuildFormatSelectorCallback(saveCallback)(analysis)
			}
		}

		application.DownloadWithMonitorAndFormatSelector(tasks, formatCallback)
	})

	// 主菜单循环
	for {
		showMainMenu()
		choice := ReadInput(i18n.T("prompt_choose_operation"))

		switch strings.TrimSpace(choice) {
		case "1":
			StartDownload(application)
		case "2":
			ModifyConfig(application)
		case "3":
			ShowAbout()
		case "4":
			ChangeLanguage(application)
		case "0":
			utils.Info(i18n.T("goodbye"))
			return
		default:
			utils.Warning(i18n.T("invalid_choice"))
		}
	}
}

// showMainMenu 显示主菜单
func showMainMenu() {
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
