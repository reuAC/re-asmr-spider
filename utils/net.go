package utils

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	proxyURL *url.URL
	proxyMux sync.RWMutex
)

// SetProxy 设置代理
func SetProxy(proxyStr string) error {
	proxyMux.Lock()
	defer proxyMux.Unlock()

	if proxyStr == "" {
		proxyURL = nil
		return nil
	}

	parsedURL, err := url.Parse(proxyStr)
	if err != nil {
		return err
	}
	proxyURL = parsedURL
	return nil
}

// getProxyFunc 获取代理函数
func getProxyFunc() func(*http.Request) (*url.URL, error) {
	proxyMux.RLock()
	defer proxyMux.RUnlock()

	if proxyURL != nil {
		return http.ProxyURL(proxyURL)
	}
	return http.ProxyFromEnvironment
}

var Client = sync.Pool{
	New: func() interface{} {
		return &http.Client{
			Timeout: 180 * time.Second, // 3分钟总超时
			Transport: &http.Transport{
				Proxy: getProxyFunc(),
				TLSClientConfig: &tls.Config{
					MaxVersion: tls.VersionTLS12, // Cloudflare 会杀
				},
				MaxIdleConns:          100,               // 最大空闲连接数
				MaxIdleConnsPerHost:   20,                // 每个host的最大空闲连接数
				IdleConnTimeout:       90 * time.Second,  // 空闲连接超时
				TLSHandshakeTimeout:   10 * time.Second,  // TLS握手超时
				ResponseHeaderTimeout: 30 * time.Second,  // 响应头超时
				ExpectContinueTimeout: 1 * time.Second,   // Expect: 100-continue超时
				DisableKeepAlives:     false,             // 启用Keep-Alive
			},
		}
	},
}
