package app

import (
	"re-asmr-spider/config"
	"re-asmr-spider/spider"
)

// App 应用主控制器
type App struct {
	Config *config.Config
}

// New 创建应用实例
func New() *App {
	return &App{
		Config: spider.Conf,
	}
}

// GetConfig 获取配置
func (a *App) GetConfig() *config.Config {
	return a.Config
}

// SaveConfig 保存配置
func (a *App) SaveConfig() error {
	return config.SaveConfig(a.Config)
}
