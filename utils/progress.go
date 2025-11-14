package utils

import (
	"fmt"
	"io"

	"github.com/schollz/progressbar/v3"
)

type ProgressBar struct {
	bar      *progressbar.ProgressBar
	finished bool
}

func NewProgressBar(total int64, prefix string) *ProgressBar {
	// 创建进度条，配置为适合文件下载的样式
	bar := progressbar.NewOptions64(
		total,
		progressbar.OptionSetDescription(prefix),
		progressbar.OptionSetWriter(io.Discard), // 先禁用输出
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(40),
		progressbar.OptionThrottle(100), // 100ms更新一次
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Println() // 完成后换行
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// 重新启用输出到 stderr (标准错误输出)
	bar.RenderBlank()
	bar = progressbar.NewOptions64(
		total,
		progressbar.OptionSetDescription(prefix),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(40),
		progressbar.OptionThrottle(100),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Println()
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	return &ProgressBar{
		bar:      bar,
		finished: false,
	}
}

func (pb *ProgressBar) Add(n int64) {
	if pb.finished {
		return
	}
	_ = pb.bar.Add64(n)
}

func (pb *ProgressBar) Finish() {
	if pb.finished {
		return
	}
	pb.finished = true
	_ = pb.bar.Finish()
}
