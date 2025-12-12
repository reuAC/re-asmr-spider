package cli

import (
	"flag"
)

// Flags 命令行参数
type Flags struct {
	ConfigPath      string
	DownloadRJs     string
	ShowVersion     bool
	ShowHelp        bool
	Account         string
	Password        string
	MaxTask         int
	MaxThread       int
	MaxRetry        int
	BufferSizeMB    int
	Proxy           string
	FormatPriority  string // 格式优先级 (例: flac,wav,mp3)
	IncludeFormats  string // 只下载指定的扩展名 (例: mp3,flac)，无论是否冲突
}

// ParseFlags 解析命令行参数
func ParseFlags() *Flags {
	f := &Flags{}

	flag.StringVar(&f.ConfigPath, "config", "", "配置文件路径 (默认: config.json)")
	flag.StringVar(&f.DownloadRJs, "download", "", "要下载的RJ号，多个用逗号分隔 (例: RJ123456,RJ789012)")
	flag.BoolVar(&f.ShowVersion, "version", false, "显示版本信息")
	flag.BoolVar(&f.ShowHelp, "help", false, "显示帮助信息")
	flag.StringVar(&f.Account, "account", "", "ASMR.one账号")
	flag.StringVar(&f.Password, "password", "", "ASMR.one密码")
	flag.IntVar(&f.MaxTask, "max-task", 0, "最大并发任务数 (默认使用配置文件)")
	flag.IntVar(&f.MaxThread, "max-thread", 0, "单文件下载线程数 (默认使用配置文件)")
	flag.IntVar(&f.MaxRetry, "max-retry", 0, "最大重试次数 (默认使用配置文件)")
	flag.IntVar(&f.BufferSizeMB, "buffer-size", 0, "缓冲区大小(MB) (默认使用配置文件)")
	flag.StringVar(&f.Proxy, "proxy", "", "HTTP/HTTPS代理地址 (例: http://127.0.0.1:7890)")
	flag.StringVar(&f.FormatPriority, "format-priority", "", "格式优先级，逗号分隔 (例: flac,wav,mp3)")
	flag.StringVar(&f.IncludeFormats, "include-formats", "", "只下载指定扩展名的文件，逗号分隔 (例: mp3,flac)，无论是否冲突")

	flag.Parse()

	return f
}

// IsCLIMode 是否为CLI模式
func (f *Flags) IsCLIMode() bool {
	return f.DownloadRJs != ""
}

// IsVersionMode 是否显示版本
func (f *Flags) IsVersionMode() bool {
	return f.ShowVersion
}

// IsHelpMode 是否显示帮助
func (f *Flags) IsHelpMode() bool {
	return f.ShowHelp
}
