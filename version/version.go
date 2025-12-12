package version

// 编译时信息
// 这些常量可以通过 go build -ldflags 在编译时注入
var (
	// AppName 应用程序名称
	AppName = "re-ASMR-Spider"
	// Version 应用程序版本
	Version = "1.1.0"
	// BuildTime 构建时间（可选，通过 ldflags 注入）
	BuildTime = "unknown"
	// GitCommit Git 提交哈希（可选，通过 ldflags 注入）
	GitCommit = "unknown"
)

// GetFullVersion 获取完整的版本信息
func GetFullVersion() string {
	return Version
}

// GetBuildInfo 获取构建信息
func GetBuildInfo() map[string]string {
	return map[string]string{
		"appName":   AppName,
		"version":   Version,
		"buildTime": BuildTime,
		"gitCommit": GitCommit,
	}
}
