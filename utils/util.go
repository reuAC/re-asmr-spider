package utils

import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || errors.Is(err, os.ErrExist)
}

// GetFileSize 获取本地文件大小
func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetRemoteFileSize 获取远程文件大小
func GetRemoteFileSize(url string, headers map[string]string) (int64, error) {
	client := Client.Get().(*http.Client)
	defer Client.Put(client)

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if _, ok := headers["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return 0, errors.New("failed to get remote file size, status: " + strconv.Itoa(resp.StatusCode))
	}

	return resp.ContentLength, nil
}
