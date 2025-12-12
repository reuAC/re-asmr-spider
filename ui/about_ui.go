package ui

import (
	"bufio"
	"os"
	"strings"

	"re-asmr-spider/i18n"
	"re-asmr-spider/utils"
)

// ShowAbout 关于页面
func ShowAbout() {
	utils.Info(i18n.T("about_title"))

	// 输出关于内容
	appName := i18n.AppName()
	version := i18n.Version()
	for _, line := range i18n.TList("about_content") {
		// 替换占位符
		line = strings.ReplaceAll(line, "{{app_name}}", appName)
		line = strings.ReplaceAll(line, "{{version}}", version)
		utils.Info(line)
	}

	utils.Info(i18n.T("press_enter_to_return"))
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
