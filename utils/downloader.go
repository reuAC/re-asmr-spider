package utils

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)


var (
	defaultUA                    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	ErrUnsupportedMultiThreading = errors.New("unsupported multi-threading")
)

type BlockMetaData struct {
	BeginOffset    int64
	EndOffset      int64
	DownloadedSize int64
}

type MultiThreadDownloader struct {
	Url         string
	SavePath    string
	FileName    string
	FullPath    string
	Client      *http.Client
	Headers     map[string]string
	Blocks      []*BlockMetaData
	ThreadCount int
	BufferSize  int // 缓冲区大小（字节），从配置文件读取
	ProgressBar *ProgressBar
	OnFailure   func(url, savePath, fileName string, err error)
	RetryCount  int
}

// progressWriter 封装 io.Writer 以更新进度条
type progressWriter struct {
	w   io.Writer
	bar *ProgressBar
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n, err := pw.w.Write(p)
	if n > 0 && pw.bar != nil {
		pw.bar.Add(int64(n))
	}
	return n, err
}

func NewDownloader(url string, path string, name string, threadCount int, bufferSize int, headers map[string]string) *MultiThreadDownloader {
	// 如果bufferSize为0，使用默认值8MB
	if bufferSize <= 0 {
		bufferSize = 8 * 1024 * 1024
	}
	return &MultiThreadDownloader{
		Url:         url,
		SavePath:    path,
		FileName:    name,
		FullPath:    path + "/" + name,
		Client:      Client.Get().(*http.Client),
		Headers:     headers,
		Blocks:      nil,
		ThreadCount: threadCount,
		BufferSize:  bufferSize,
	}
}

func (m *MultiThreadDownloader) Download() error {
	if m.ThreadCount < 2 {
		err := m.singleThreadDownload()
		return err
	}
	if err := m.initDownload(); err != nil {
		if err == ErrUnsupportedMultiThreading {
			return nil
		}
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(len(m.Blocks))
	var lastErr error
	for i := range m.Blocks {
		go func(b *BlockMetaData) {
			defer wg.Done()
			if err := m.downloadBlocks(b); err != nil {
				lastErr = err
			}
		}(m.Blocks[i])
	}
	wg.Wait()
	if m.ProgressBar != nil {
		m.ProgressBar.Finish()
	}
	return lastErr
}

func (m *MultiThreadDownloader) initDownload() error {
	var contentLength int64

	// 辅助函数：使用 io.CopyBuffer 优化流式复制
	copyStream := func(s io.ReadCloser, size int64) error {
		file, err := os.OpenFile(m.FullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			return err
		}
		defer file.Close()

		// 使用 bufio 减少磁盘 IO 系统调用
		writer := bufio.NewWriterSize(file, m.BufferSize)
		defer writer.Flush()

		// 创建进度条
		if size > 0 {
			m.ProgressBar = NewProgressBar(size, m.FileName)
		}

		// 封装 writer 以自动更新进度
		pw := &progressWriter{w: writer, bar: m.ProgressBar}

		buf := make([]byte, m.BufferSize)
		_, err = io.CopyBuffer(pw, s, buf)
		if err != nil {
			return err
		}

		if m.ProgressBar != nil {
			m.ProgressBar.Finish()
		}
		return ErrUnsupportedMultiThreading // 按照原逻辑返回此错误以终止后续多线程逻辑
	}
	req, err := http.NewRequest("GET", m.Url, nil)
	if err != nil {
		return err
	}

	for k, v := range m.Headers {
		req.Header.Set(k, v)
	}
	if _, ok := m.Headers["User-Agent"]; !ok {
		req.Header["User-Agent"] = []string{defaultUA}
	}
	// 尝试获取文件头信息或探测 Range 支持
	req.Header.Set("range", "bytes=0-")
	resp, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("response status unsuccessful: " + strconv.FormatInt(int64(resp.StatusCode), 10))
	}
	// 如果服务器直接返回 200 (不支持 Range) 或者没有 ContentLength
	if resp.StatusCode == 200 {
		return copyStream(resp.Body, resp.ContentLength)
	}

	if resp.StatusCode == 206 {
		contentLength = resp.ContentLength
		// 创建进度条
		if contentLength > 0 {
			m.ProgressBar = NewProgressBar(contentLength, m.FileName)
		}

		blockSize := func() int64 {
			if contentLength > 1024*1024 {
				return (contentLength / int64(m.ThreadCount)) - 10
			}
			return contentLength
		}()

		// 如果块大小等于内容长度，说明不需要分块
		if blockSize == contentLength {
			return copyStream(resp.Body, contentLength)
		}

		// 计算分块
		var tmp int64
		for tmp+blockSize < contentLength {
			m.Blocks = append(m.Blocks, &BlockMetaData{
				BeginOffset: tmp,
				EndOffset:   tmp + blockSize - 1,
			})
			tmp += blockSize
		}
		m.Blocks = append(m.Blocks, &BlockMetaData{
			BeginOffset: tmp,
			EndOffset:   contentLength - 1,
		})
		return nil
	}
	return errors.New("unknown status code")
}

func (m *MultiThreadDownloader) downloadBlocks(block *BlockMetaData) error {
	req, _ := http.NewRequest("GET", m.Url, nil)
	file, err := os.OpenFile(m.FullPath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 定位到该块的起始位置
	if _, err := file.Seek(block.BeginOffset, io.SeekStart); err != nil {
		return err
	}

	// 使用配置的BufferSize来优化IO性能
	writer := bufio.NewWriterSize(file, m.BufferSize)
	defer writer.Flush()

	for k, v := range m.Headers {
		req.Header.Set(k, v)
	}
	if _, ok := m.Headers["User-Agent"]; !ok {
		req.Header["User-Agent"] = []string{defaultUA}
	}
	req.Header.Set("range", "bytes="+strconv.FormatInt(block.BeginOffset, 10)+"-"+strconv.FormatInt(block.EndOffset, 10))

	resp, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("response status unsuccessful: " + strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	// 使用大buffer来减少系统调用，提升VPS等网络环境下的性能
	buffer := make([]byte, m.BufferSize)

	for {
		n, readErr := resp.Body.Read(buffer)
		if n > 0 {
			// 计算需要写入的大小，防止多写
			bytesToWrite := int64(n)
			remaining := block.EndOffset + 1 - block.BeginOffset
			if bytesToWrite > remaining {
				bytesToWrite = remaining
			}

			if _, writeErr := writer.Write(buffer[:bytesToWrite]); writeErr != nil {
				return writeErr
			}

			block.BeginOffset += bytesToWrite
			block.DownloadedSize += bytesToWrite

			if m.ProgressBar != nil {
				m.ProgressBar.Add(bytesToWrite)
			}

			if block.BeginOffset > block.EndOffset {
				break
			}
		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return readErr
		}
	}
	return nil
}

func (m *MultiThreadDownloader) singleThreadDownload() error {
	file, err := os.OpenFile(m.FullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 使用 bufio 优化磁盘写入
	writer := bufio.NewWriterSize(file, m.BufferSize)
	defer writer.Flush()

	req, err := http.NewRequest("GET", m.Url, nil)
	if err != nil {
		return err
	}

	for k, v := range m.Headers {
		req.Header.Set(k, v)
	}
	if _, ok := m.Headers["User-Agent"]; !ok {
		req.Header["User-Agent"] = []string{defaultUA}
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建进度条
	if resp.ContentLength > 0 {
		m.ProgressBar = NewProgressBar(resp.ContentLength, m.FileName)
	}

	// 使用 io.CopyBuffer 替代手动循环，提升性能
	pw := &progressWriter{w: writer, bar: m.ProgressBar}
	buf := make([]byte, m.BufferSize)

	if _, err := io.CopyBuffer(pw, resp.Body, buf); err != nil {
		return err
	}

	if m.ProgressBar != nil {
		m.ProgressBar.Finish()
	}
	return nil
}
