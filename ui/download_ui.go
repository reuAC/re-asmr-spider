package ui

import (
	"strings"

	"re-asmr-spider/app"
	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

// StartDownload 开始下载界面
func StartDownload(application *app.App) {
	utils.Info(i18n.T("start_download_title"))
	utils.Info(i18n.T("download_input_hint"))
	utils.Info(i18n.T("download_multiple_hint"))

	rjNumbers := ReadInput(i18n.T("prompt_rj_number"))
	if rjNumbers == "" {
		utils.Warning(i18n.T("no_rj_input"))
		return
	}

	tasks := strings.Split(rjNumbers, " ")
	// BuildFormatSelectorCallback 会接收保存回调作为参数
	application.DownloadWithMonitorAndFormatSelector(tasks, func(analysis *spider.FormatAnalysis, saveCallback func(*spider.FilterStrategy)) *spider.FilterStrategy {
		return BuildFormatSelectorCallback(saveCallback)(analysis)
	})
}
