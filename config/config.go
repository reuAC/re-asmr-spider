package config

import (
	"encoding/json"
	"io"
	"os"
)

type FormatStrategy struct {
	Mode             string              `json:"mode"`               // "all", "priority", "manual"
	PriorityFormats  []string            `json:"priority_formats"`   // 优先格式列表
	ManualSelections map[string][]string `json:"manual_selections"`  // 手动选择记录
	IncludeFormats   []string            `json:"include_formats"`    // 额外包含的扩展名
}

type DownloadState struct {
	InProgress      bool            `json:"in_progress"`
	Tasks           []string        `json:"tasks"`
	FormatStrategy  *FormatStrategy `json:"format_strategy,omitempty"` // 格式筛选策略
}

type Config struct {
	Account       string        `json:"account"`
	Password      string        `json:"password"`
	MaxTask       int           `json:"max_task"`
	MaxThread     int           `json:"max_thread"`
	MaxRetry      int           `json:"max_retry"`
	Language      string        `json:"language"`
	Proxy         string        `json:"proxy"`
	BufferSizeMB  int           `json:"buffer_size_mb"`
	DownloadState DownloadState `json:"download_state"`
}

func generateDefaultConfig() *Config {
	return &Config{
		Account:      "guest",
		Password:     "guest",
		MaxTask:      1,
		MaxThread:    1,
		MaxRetry:     3,
		Language:     "zh-CN",
		Proxy:        "",
		BufferSizeMB: 8, // 默认8MB，适用于VPS和云存储环境
		DownloadState: DownloadState{
			InProgress: false,
			Tasks:      []string{},
		},
	}
}

func SaveConfig(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("config.json", data, 0644)
}

// SaveDownloadState 保存下载状态（不含格式策略）
func SaveDownloadState(cfg *Config, tasks []string) error {
	cfg.DownloadState.InProgress = true
	cfg.DownloadState.Tasks = tasks
	return SaveConfig(cfg)
}

// SaveDownloadStateWithStrategy 保存下载状态（含格式策略）
func SaveDownloadStateWithStrategy(cfg *Config, tasks []string, strategy *FormatStrategy) error {
	cfg.DownloadState.InProgress = true
	cfg.DownloadState.Tasks = tasks
	cfg.DownloadState.FormatStrategy = strategy
	return SaveConfig(cfg)
}

// ClearDownloadState 清除下载状态
func ClearDownloadState(cfg *Config) error {
	cfg.DownloadState.InProgress = false
	cfg.DownloadState.Tasks = []string{}
	cfg.DownloadState.FormatStrategy = nil
	return SaveConfig(cfg)
}

func GetConfig() (*Config, error) {
	return GetConfigFromPath("config.json")
}

// GetConfigFromPath 从指定路径加载配置文件
func GetConfigFromPath(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 配置文件不存在，创建默认配置
		cfg := generateDefaultConfig()
		if saveErr := SaveConfig(cfg); saveErr != nil {
			return nil, saveErr
		}
		return cfg, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	all, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(all, &config)
	if err != nil {
		return nil, err
	}

	// 验证和修正配置值
	validateConfig(&config)

	return &config, nil
}

// validateConfig 验证并修正配置值
func validateConfig(cfg *Config) {
	// 验证BufferSizeMB，如果为0或负数，设置为默认值8MB
	if cfg.BufferSizeMB <= 0 {
		cfg.BufferSizeMB = 8
	}
	// 限制最小值为1MB，最大值为64MB
	if cfg.BufferSizeMB < 1 {
		cfg.BufferSizeMB = 1
	}
	if cfg.BufferSizeMB > 64 {
		cfg.BufferSizeMB = 64
	}
}

// GetBufferSize 获取字节单位的buffer大小
func (c *Config) GetBufferSize() int {
	return c.BufferSizeMB * 1024 * 1024
}
