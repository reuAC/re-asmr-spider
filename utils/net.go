package utils

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	proxyURL     *url.URL
	proxyMux     sync.RWMutex
	globalClient *http.Client
	clientOnce   sync.Once
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

// dynamicProxyFunc 动态代理函数，每次请求时都会读取最新的代理设置
func dynamicProxyFunc(req *http.Request) (*url.URL, error) {
	proxyMux.RLock()
	defer proxyMux.RUnlock()

	if proxyURL != nil {
		return proxyURL, nil
	}
	return http.ProxyFromEnvironment(req)
}

// initGlobalClient 初始化全局HTTP客户端（单例模式）
func initGlobalClient() *http.Client {
	clientOnce.Do(func() {
		globalClient = &http.Client{
			Timeout: 180 * time.Second, // 3分钟总超时
			Transport: &http.Transport{
				Proxy: dynamicProxyFunc, // 使用动态代理函数
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
	})
	return globalClient
}

// ClientPool 提供兼容原有sync.Pool接口的Client管理
type ClientPool struct{}

var Client = &ClientPool{}

// Get 获取HTTP客户端（兼容原有接口）
func (cp *ClientPool) Get() interface{} {
	return initGlobalClient()
}

// Put 归还HTTP客户端（兼容原有接口，实际不做任何操作）
func (cp *ClientPool) Put(interface{}) {
	// 使用全局单例，不需要归还操作
}
