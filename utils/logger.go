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

// printLog 内部方法：打印日志消息
func (l *Logger) printLog(level LogLevel, message string) {
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

	if l.enableColor {
		if l.showTime {
			timestamp := time.Now().Format("15:04:05")
			fmt.Print(colorGray, timestamp, colorReset, " ", color, prefix, colorReset, " ", message, "\n")
		} else {
			fmt.Print(color, prefix, colorReset, " ", message, "\n")
		}
	} else {
		if l.showTime {
			timestamp := time.Now().Format("15:04:05")
			fmt.Print(timestamp, " ", prefix, " ", message, "\n")
		} else {
			fmt.Print(prefix, " ", message, "\n")
		}
	}
}

// logMessage 内部方法：输出简单消息（不格式化）
func (l *Logger) logMessage(level LogLevel, msg string) {
	l.printLog(level, msg)
}

// logFormat 内部方法：格式化输出
func (l *Logger) logFormat(level LogLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.printLog(level, message)
}

// Debug 调试日志（灰色）- 输出消息
func Debug(msg string) {
	defaultLogger.logMessage(LevelDebug, msg)
}

// Debugf 调试日志（灰色）- 格式化输出
func Debugf(format string, args ...interface{}) {
	defaultLogger.logFormat(LevelDebug, format, args...)
}

// Info 信息日志（蓝色）- 输出消息
func Info(msg string) {
	defaultLogger.logMessage(LevelInfo, msg)
}

// Infof 信息日志（蓝色）- 格式化输出
func Infof(format string, args ...interface{}) {
	defaultLogger.logFormat(LevelInfo, format, args...)
}

// Success 成功日志（绿色）- 输出消息
func Success(msg string) {
	defaultLogger.logMessage(LevelSuccess, msg)
}

// Successf 成功日志（绿色）- 格式化输出
func Successf(format string, args ...interface{}) {
	defaultLogger.logFormat(LevelSuccess, format, args...)
}

// Warning 警告日志（黄色）- 输出消息
func Warning(msg string) {
	defaultLogger.logMessage(LevelWarning, msg)
}

// Warningf 警告日志（黄色）- 格式化输出
func Warningf(format string, args ...interface{}) {
	defaultLogger.logFormat(LevelWarning, format, args...)
}

// Error 错误日志（红色）- 输出消息
func Error(msg string) {
	defaultLogger.logMessage(LevelError, msg)
}

// Errorf 错误日志（红色）- 格式化输出
func Errorf(format string, args ...interface{}) {
	defaultLogger.logFormat(LevelError, format, args...)
}

// SetColorEnabled 设置是否启用颜色
func SetColorEnabled(enabled bool) {
	defaultLogger.enableColor = enabled
}

// SetShowTime 设置是否显示时间
func SetShowTime(show bool) {
	defaultLogger.showTime = show
}
