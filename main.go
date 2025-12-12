package main

import (
	"fmt"
	"os"

	"re-asmr-spider/app"
	"re-asmr-spider/cli"
	"re-asmr-spider/config"
	"re-asmr-spider/spider"
	"re-asmr-spider/ui"
	"re-asmr-spider/version"
)

func main() {
	// 解析命令行参数
	flags := cli.ParseFlags()

	// 处理版本信息（不需要配置）
	if flags.IsVersionMode() {
		// 使用默认应用名称
		fmt.Printf("ASMR Downloader v%s\n", version.Version)
		return
	}

	// 处理帮助信息（不需要配置）
	if flags.IsHelpMode() {
		cli.PrintHelp()
		return
	}

	// 初始化配置
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化spider包（包括i18n和代理）
	if err := spider.InitSpider(cfg); err != nil {
		fmt.Printf("Failed to initialize: %v\n", err)
		os.Exit(1)
	}

	// 创建应用实例
	application := app.New()

	// CLI模式
	if flags.IsCLIMode() {
		command := cli.NewCommand(flags, application)
		command.ApplyFlags()
		command.Execute()
		return
	}

	// 交互模式
	ui.RunInteractive(application)
}
