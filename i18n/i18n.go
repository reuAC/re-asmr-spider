package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"sync"

	"re-asmr-spider/version"
)

//go:embed locales/*.json global.json
var localesFS embed.FS

type Language struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type GlobalConfig struct {
	Languages []Language `json:"languages"`
}

var (
	currentLocale   = "zh-CN"
	translations    = make(map[string]map[string]interface{})
	translationsMux sync.RWMutex
	globalConfig    GlobalConfig
)

// Init 初始化i18n模块
func Init(locale string) error {
	translationsMux.Lock()
	defer translationsMux.Unlock()

	// 加载全局配置
	globalData, err := localesFS.ReadFile("global.json")
	if err != nil {
		return fmt.Errorf("failed to read global.json: %w", err)
	}
	if err := json.Unmarshal(globalData, &globalConfig); err != nil {
		return fmt.Errorf("failed to parse global.json: %w", err)
	}

	// 加载所有支持的语言
	for _, lang := range globalConfig.Languages {
		data, err := localesFS.ReadFile("locales/" + lang.Code + ".json")
		if err != nil {
			// 如果文件不存在，跳过（可能还没有翻译）
			continue
		}

		var trans map[string]interface{}
		if err := json.Unmarshal(data, &trans); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", lang.Code, err)
		}

		translations[lang.Code] = trans
	}

	// 设置当前语言
	if locale != "" {
		currentLocale = locale
	}

	return nil
}

// SetLocale 设置当前语言
func SetLocale(locale string) {
	translationsMux.Lock()
	defer translationsMux.Unlock()

	if _, ok := translations[locale]; ok {
		currentLocale = locale
	}
}

// GetLocale 获取当前语言
func GetLocale() string {
	translationsMux.RLock()
	defer translationsMux.RUnlock()
	return currentLocale
}

// T 获取翻译文本（支持格式化参数）
func T(key string, args ...interface{}) string {
	translationsMux.RLock()
	defer translationsMux.RUnlock()

	trans, ok := translations[currentLocale]
	if !ok {
		// 回退到中文
		trans = translations["zh-CN"]
	}

	value, ok := trans[key]
	if !ok {
		return key // 如果找不到翻译，返回key本身
	}

	// 处理字符串类型
	if str, ok := value.(string); ok {
		if len(args) > 0 {
			return fmt.Sprintf(str, args...)
		}
		return str
	}

	// 处理数组类型（返回第一个元素）
	if arr, ok := value.([]interface{}); ok && len(arr) > 0 {
		if str, ok := arr[0].(string); ok {
			if len(args) > 0 {
				return fmt.Sprintf(str, args...)
			}
			return str
		}
	}

	return key
}

// TList 获取翻译文本列表
func TList(key string) []string {
	translationsMux.RLock()
	defer translationsMux.RUnlock()

	trans, ok := translations[currentLocale]
	if !ok {
		// 回退到中文
		trans = translations["zh-CN"]
	}

	value, ok := trans[key]
	if !ok {
		return []string{key}
	}

	// 处理数组类型
	if arr, ok := value.([]interface{}); ok {
		result := make([]string, 0, len(arr))
		for _, item := range arr {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	}

	// 处理字符串类型（包装成数组）
	if str, ok := value.(string); ok {
		return []string{str}
	}

	return []string{key}
}

// GetSupportedLocales 获取支持的语言列表
func GetSupportedLocales() []Language {
	return globalConfig.Languages
}

// GetLocaleName 获取语言显示名称
func GetLocaleName(locale string) string {
	for _, lang := range globalConfig.Languages {
		if lang.Code == locale {
			return lang.NativeName
		}
	}
	return locale
}

// GetLanguageByCode 根据语言代码获取语言信息
func GetLanguageByCode(code string) *Language {
	for _, lang := range globalConfig.Languages {
		if lang.Code == code {
			return &lang
		}
	}
	return nil
}

// AppName 获取应用程序名称（统一从 version 包获取）
func AppName() string {
	return version.AppName
}

// Version 获取应用程序版本（统一从 version 包获取）
func Version() string {
	return version.Version
}
