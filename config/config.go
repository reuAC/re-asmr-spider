package config

import (
	"encoding/json"
	"io"
	"os"
)

type DownloadState struct {
	InProgress bool     `json:"in_progress"`
	Tasks      []string `json:"tasks"`
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

// SaveDownloadState 保存下载状态
func SaveDownloadState(cfg *Config, tasks []string) error {
	cfg.DownloadState.InProgress = true
	cfg.DownloadState.Tasks = tasks
	return SaveConfig(cfg)
}

// ClearDownloadState 清除下载状态
func ClearDownloadState(cfg *Config) error {
	cfg.DownloadState.InProgress = false
	cfg.DownloadState.Tasks = []string{}
	return SaveConfig(cfg)
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		// 配置文件不存在，创建默认配置
		cfg := generateDefaultConfig()
		if saveErr := SaveConfig(cfg); saveErr != nil {
			return nil, saveErr
		}
		return cfg, nil
	}

	file, err := os.Open("config.json")
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
