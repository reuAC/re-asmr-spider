package utils

import (
	"sync"
	"time"
)

// ActivityMonitor 活动监控器
type ActivityMonitor struct {
	lastActivity time.Time
	mu           sync.RWMutex
	enabled      bool
	timeout      time.Duration
}

var GlobalMonitor = &ActivityMonitor{
	lastActivity: time.Now(),
	enabled:      false,
	timeout:      5 * time.Minute, // 默认5分钟无活动超时
}

// Start 启动监控
func (am *ActivityMonitor) Start() {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.enabled = true
	am.lastActivity = time.Now()
}

// Stop 停止监控
func (am *ActivityMonitor) Stop() {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.enabled = false
}

// UpdateActivity 更新活动时间
func (am *ActivityMonitor) UpdateActivity() {
	am.mu.Lock()
	defer am.mu.Unlock()
	if am.enabled {
		am.lastActivity = time.Now()
	}
}

// GetInactiveTime 获取无活动时间
func (am *ActivityMonitor) GetInactiveTime() time.Duration {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return time.Since(am.lastActivity)
}

// IsTimeout 检查是否超时
func (am *ActivityMonitor) IsTimeout() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	if !am.enabled {
		return false
	}
	return time.Since(am.lastActivity) > am.timeout
}

// IsEnabled 检查监控是否启用
func (am *ActivityMonitor) IsEnabled() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.enabled
}

// SetTimeout 设置超时时间
func (am *ActivityMonitor) SetTimeout(timeout time.Duration) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.timeout = timeout
}
