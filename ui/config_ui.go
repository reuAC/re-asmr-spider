package ui

import (
	"strconv"
	"strings"

	"re-asmr-spider/app"
	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/utils"
)

// ModifyConfig 配置修改界面
func ModifyConfig(application *app.App) {
	utils.Info(i18n.T("modify_config_title"))

	cfg := application.GetConfig()

	// 显示当前配置
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
	utils.Info("  " + i18n.T("config_buffer_size", cfg.BufferSizeMB))
	utils.Info("  " + i18n.T("config_return"))

	choice := ReadInput("\n" + i18n.T("prompt_choose_modify"))

	switch choice {
	case "1":
		newAccount := ReadInput(i18n.T("prompt_new_account"))
		if newAccount != "" {
			cfg.Account = newAccount
			utils.Success(i18n.T("account_updated"))
		}
	case "2":
		newPassword := ReadInput(i18n.T("prompt_new_password"))
		if newPassword != "" {
			cfg.Password = newPassword
			utils.Success(i18n.T("password_updated"))
		}
	case "3":
		newMaxTask := ReadInput(i18n.T("prompt_new_max_task"))
		if val, err := strconv.Atoi(newMaxTask); err == nil && val > 0 {
			cfg.MaxTask = val
			utils.Success(i18n.T("max_task_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "4":
		newMaxThread := ReadInput(i18n.T("prompt_new_max_thread"))
		if val, err := strconv.Atoi(newMaxThread); err == nil && val > 0 {
			cfg.MaxThread = val
			utils.Success(i18n.T("max_thread_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "5":
		newMaxRetry := ReadInput(i18n.T("prompt_new_max_retry"))
		if val, err := strconv.Atoi(newMaxRetry); err == nil && val >= 0 {
			cfg.MaxRetry = val
			utils.Success(i18n.T("max_retry_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_value"))
		}
	case "6":
		ChangeLanguage(application)
		return
	case "7":
		newProxy := ReadInput(i18n.T("prompt_new_proxy"))
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
	case "8":
		newBufferSize := ReadInput(i18n.T("prompt_new_buffer_size"))
		if val, err := strconv.Atoi(newBufferSize); err == nil && val >= 1 && val <= 64 {
			cfg.BufferSizeMB = val
			utils.Success(i18n.T("buffer_size_updated", val))
		} else {
			utils.Warning(i18n.T("invalid_buffer_size"))
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
