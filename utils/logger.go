package utils

import (
	"fmt"
	"runtime"
	"time"
)

// ANSI 颜色代码
const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorGray    = "\033[90m"
	colorBold    = "\033[1m"
)

// 日志级别
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelSuccess
	LevelWarning
	LevelError
)

// Logger 日志记录器
type Logger struct {
	enableColor bool
	showTime    bool
}

var defaultLogger = &Logger{
	enableColor: runtime.GOOS != "windows" || isWindowsColorSupported(),
	showTime:    false,
}

// isWindowsColorSupported 检查Windows是否支持ANSI颜色（Windows 10+）
func isWindowsColorSupported() bool {
	// Windows 10+ 默认支持ANSI转义序列
	return true
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	var prefix string
	var color string

	switch level {
	case LevelDebug:
		prefix = "[DEBUG]"
		color = colorGray
	case LevelInfo:
		prefix = "[INFO]"
		color = colorBlue
	case LevelSuccess:
		prefix = "[SUCCESS]"
		color = colorGreen
	case LevelWarning:
		prefix = "[WARN]"
		color = colorYellow
	case LevelError:
		prefix = "[ERROR]"
		color = colorRed
	}

	message := fmt.Sprintf(format, args...)

	if l.enableColor {
		if l.showTime {
			timestamp := time.Now().Format("15:04:05")
			fmt.Printf("%s%s%s %s%s%s %s\n", colorGray, timestamp, colorReset, color, prefix, colorReset, message)
		} else {
			fmt.Printf("%s%s%s %s\n", color, prefix, colorReset, message)
		}
	} else {
		if l.showTime {
			timestamp := time.Now().Format("15:04:05")
			fmt.Printf("%s %s %s\n", timestamp, prefix, message)
		} else {
			fmt.Printf("%s %s\n", prefix, message)
		}
	}
}

// Debug 调试日志（灰色）
func Debug(format string, args ...interface{}) {
	defaultLogger.log(LevelDebug, format, args...)
}

// Info 信息日志（蓝色）
func Info(format string, args ...interface{}) {
	defaultLogger.log(LevelInfo, format, args...)
}

// Success 成功日志（绿色）
func Success(format string, args ...interface{}) {
	defaultLogger.log(LevelSuccess, format, args...)
}

// Warning 警告日志（黄色）
func Warning(format string, args ...interface{}) {
	defaultLogger.log(LevelWarning, format, args...)
}

// Error 错误日志（红色）
func Error(format string, args ...interface{}) {
	defaultLogger.log(LevelError, format, args...)
}

// SetColorEnabled 设置是否启用颜色
func SetColorEnabled(enabled bool) {
	defaultLogger.enableColor = enabled
}

// SetShowTime 设置是否显示时间
func SetShowTime(show bool) {
	defaultLogger.showTime = show
}
