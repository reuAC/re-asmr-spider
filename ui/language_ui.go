package ui

import (
	"strconv"
	"strings"

	"re-asmr-spider/app"
	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/utils"
)

// ChangeLanguage 语言切换界面
func ChangeLanguage(application *app.App) {
	utils.Info(i18n.T("language_title"))

	// 动态显示所有支持的语言
	languages := i18n.GetSupportedLocales()
	langMap := make(map[string]string)

	for idx, lang := range languages {
		num := idx + 1
		utils.Infof("  %d. %s", num, lang.NativeName)
		langMap[strconv.Itoa(num)] = lang.Code
	}
	utils.Info("  " + i18n.T("language_return"))

	choice := ReadInput(i18n.T("prompt_choose_language"))
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
	cfg := application.GetConfig()
	cfg.Language = newLang
	i18n.SetLocale(newLang)

	// 保存配置
	if err := config.SaveConfig(cfg); err != nil {
		utils.Error(i18n.T("save_config_failed", err))
	} else {
		utils.Success(i18n.T("language_updated", i18n.GetLocaleName(newLang)))
	}
}
