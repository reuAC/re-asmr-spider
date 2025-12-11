package spider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"

	"re-asmr-spider/config"
	"re-asmr-spider/i18n"
	"re-asmr-spider/utils"
)

var Conf *config.Config

func init() {
	var err error
	Conf, err = config.GetConfig()
	if err != nil {
		fmt.Printf("Failed to get config: %v\n", err)
		os.Exit(1)
	}

	// 初始化i18n (使用配置中的语言设置)
	if err := i18n.Init(Conf.Language); err != nil {
		fmt.Printf("Failed to initialize i18n: %v\n", err)
		os.Exit(1)
	}

	// 初始化代理设置
	if Conf.Proxy != "" {
		if err := utils.SetProxy(Conf.Proxy); err != nil {
			fmt.Printf("Failed to set proxy: %v\n", err)
		}
	}
}

type FailedTask struct {
	URL       string
	DirPath   string
	FileName  string
	RetryCount int
}

type ASMRClient struct {
	Authorization string
	WorkerPool    *utils.WorkerPool
	ThreadCount   int
	BufferSize    int
	FailedTasks   []FailedTask
	MaxRetry      int
	mu            sync.Mutex
}

type track struct {
	Type             string  `json:"type"`
	Title            string  `json:"title"`
	Children         []track `json:"children,omitempty"`
	Hash             string  `json:"hash,omitempty"`
	WorkTitle        string  `json:"workTitle,omitempty"`
	MediaStreamURL   string  `json:"mediaStreamUrl,omitempty"`
	MediaDownloadURL string  `json:"mediaDownloadUrl,omitempty"`
}

func NewASMRClient(maxTask int, maxThread int, maxRetry int, bufferSize int) *ASMRClient {
	return &ASMRClient{
		WorkerPool:  utils.NewWorkerPool(maxTask),
		ThreadCount: maxThread,
		BufferSize:  bufferSize,
		FailedTasks: make([]FailedTask, 0),
		MaxRetry:    maxRetry,
	}
}

func (ac *ASMRClient) Login() error {
	utils.GlobalMonitor.UpdateActivity()

	payload, err := json.Marshal(map[string]string{
		"name":     Conf.Account,
		"password": Conf.Password,
	})
	if err != nil {
		utils.Error(i18n.T("parse_error", err))
		return err
	}
	client := utils.Client.Get().(*http.Client)
	req, _ := http.NewRequest("POST", "https://api.asmr.one/api/auth/me", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://www.asmr.one/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	resp, err := client.Do(req)
	utils.Client.Put(client)
	if err != nil {
		utils.Error(i18n.T("network_error", err))
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.Error(i18n.T("request_failed", err))
		return err
	}
	res := make(map[string]string)
	err = json.Unmarshal(all, &res)
	ac.Authorization = "Bearer " + res["token"]
	utils.GlobalMonitor.UpdateActivity()
	utils.Success(i18n.T("login_success"))
	return nil
}

func (ac *ASMRClient) GetVoiceTracks(id string) ([]track, error) {
	utils.GlobalMonitor.UpdateActivity()

	client := utils.Client.Get().(*http.Client)
	req, _ := http.NewRequest("GET", "https://api.asmr.one/api/tracks/"+id, nil)
	req.Header.Set("Authorization", ac.Authorization)
	req.Header.Set("Referer", "https://www.asmr.one/")
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.0")
	resp, err := client.Do(req)
	utils.Client.Put(client)
	if err != nil {
		utils.Error(i18n.T("request_failed", err))
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.Error(i18n.T("request_failed", err))
		return nil, err
	}
	res := make([]track, 0)
	err = json.Unmarshal(all, &res)
	utils.GlobalMonitor.UpdateActivity()
	return res, nil
}

// AddFailedTask 添加失败任务到重试队列
func (ac *ASMRClient) AddFailedTask(url, dirPath, fileName string, retryCount int) {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	ac.FailedTasks = append(ac.FailedTasks, FailedTask{
		URL:       url,
		DirPath:   dirPath,
		FileName:  fileName,
		RetryCount: retryCount,
	})
}

// RetryFailedTasks 重试所有失败的任务
func (ac *ASMRClient) RetryFailedTasks() bool {
	ac.mu.Lock()
	if len(ac.FailedTasks) == 0 {
		ac.mu.Unlock()
		return false
	}

	tasks := make([]FailedTask, len(ac.FailedTasks))
	copy(tasks, ac.FailedTasks)
	ac.FailedTasks = make([]FailedTask, 0)
	ac.mu.Unlock()

	utils.Warning(i18n.T("retrying", len(tasks), ac.MaxRetry))

	permanentlyFailed := make([]FailedTask, 0)
	retriedCount := 0
	for _, task := range tasks {
		if task.RetryCount >= ac.MaxRetry {
			utils.Error(i18n.T("max_retry_reached", task.FileName))
			permanentlyFailed = append(permanentlyFailed, task)
			continue
		}
		utils.Info(i18n.T("retrying", task.RetryCount+1, ac.MaxRetry) + ": " + task.FileName)
		ac.downloadFileWithRetry(task.URL, task.DirPath, task.FileName, task.RetryCount+1)
		retriedCount++
	}

	// 将永久失败的任务放回列表
	if len(permanentlyFailed) > 0 {
		ac.mu.Lock()
		ac.FailedTasks = append(ac.FailedTasks, permanentlyFailed...)
		ac.mu.Unlock()
	}

	// 只有当有任务被重试时才返回 true
	return retriedCount > 0
}

func (ac *ASMRClient) Download(id string) {
	id = strings.Replace(id, "RJ", "", 1)
	utils.Info(i18n.T("fetching_work_info", "RJ"+id))
	tracks, err := ac.GetVoiceTracks(id)
	if err != nil {
		utils.Error(i18n.T("request_failed", err))
		return
	}
	basePath := "downloads/RJ" + id
	ac.EnsureDir(tracks, basePath)
	utils.Success(i18n.T("work_info_fetched", "RJ"+id))
}

func (ac *ASMRClient) downloadFileWithRetry(url string, dirPath string, fileName string, retryCount int) {
	ac.downloadFileInternal(url, dirPath, fileName, retryCount)
}

func (ac *ASMRClient) DownloadFile(url string, dirPath string, fileName string) {
	ac.downloadFileInternal(url, dirPath, fileName, 0)
}

func (ac *ASMRClient) downloadFileInternal(url string, dirPath string, fileName string, retryCount int) {
	if runtime.GOOS == "windows" {
		for _, str := range []string{"?", "<", ">", ":", "/", "\\", "*", "|"} {
			fileName = strings.Replace(fileName, str, "_", -1)
		}
	}
	savePath := dirPath + "/" + fileName
	headers := map[string]string{
		"Referer": "https://www.asmr.one/",
	}

	if utils.PathExists(savePath) {
		// 获取本地文件大小
		localSize, err := utils.GetFileSize(savePath)
		if err != nil {
			utils.Warning(i18n.T("file_error", err))
		} else {
			// 获取远程文件大小
			remoteSize, err := utils.GetRemoteFileSize(url, headers)
			if err != nil {
				utils.Warning(i18n.T("network_error", err))
				utils.Info(i18n.T("file_exists", savePath))
				return
			}

			if localSize == remoteSize {
				utils.Info(i18n.T("file_exists", savePath))
				return
			} else {
				utils.Warning(i18n.T("file_error", fmt.Sprintf("size mismatch: local=%d, remote=%d", localSize, remoteSize)))
			}
		}
	}

	downloader := utils.NewDownloader(url, dirPath, fileName, ac.ThreadCount, ac.BufferSize, headers)
	downloader.RetryCount = retryCount
	downloader.OnFailure = func(failedUrl, failedPath, failedName string, err error) {
		ac.AddFailedTask(failedUrl, failedPath, failedName, retryCount)
	}
	ac.WorkerPool.TaskQueue <- downloader
}

func (ac *ASMRClient) EnsureDir(tracks []track, basePath string) {
	path := basePath
	_ = os.MkdirAll(path, os.ModePerm)
	for _, t := range tracks {
		if t.Type != "folder" {
			ac.DownloadFile(t.MediaDownloadURL, path, t.Title)
		} else {
			ac.EnsureDir(t.Children, path+"/"+t.Title)
		}
	}
}
