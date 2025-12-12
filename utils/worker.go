package utils

import (
	"os"
	"sync"

	"re-asmr-spider/i18n"
)

type WorkerChan chan *MultiThreadDownloader

type WorkerPool struct {
	sync.WaitGroup
	cond      *sync.Cond
	TaskQueue WorkerChan
	Limit     int
	Count     int
	closed    bool
	closeMux  sync.Mutex
}

func NewWorkerPool(WorkerCount int) *WorkerPool {
	return &WorkerPool{
		cond:      sync.NewCond(&sync.Mutex{}),
		Limit:     WorkerCount,
		TaskQueue: make(WorkerChan, WorkerCount),
	}
}

// Submit 提交任务到工作池（在提交时就增加 WaitGroup 计数，避免竞态条件）
func (wp *WorkerPool) Submit(task *MultiThreadDownloader) {
	wp.closeMux.Lock()
	defer wp.closeMux.Unlock()

	if wp.closed {
		return // 如果已关闭，则忽略新任务
	}

	wp.Add(1)
	wp.TaskQueue <- task
}

// Close 关闭工作池，不再接受新任务
func (wp *WorkerPool) Close() {
	wp.closeMux.Lock()
	defer wp.closeMux.Unlock()

	if wp.closed {
		return // 避免重复关闭
	}

	wp.closed = true
	close(wp.TaskQueue)
}

func (wp *WorkerPool) Start() {
	go func() {
		for t := range wp.TaskQueue {
			wp.cond.L.Lock()
			for wp.Count >= wp.Limit {
				wp.cond.Wait()
			}
			wp.cond.L.Unlock()
			go func(t *MultiThreadDownloader) {
				wp.cond.L.Lock()
				wp.Count++
				wp.cond.L.Unlock()
				defer func() {
					wp.cond.L.Lock()
					wp.Count--
					wp.Done()
					wp.cond.Broadcast()
					wp.cond.L.Unlock()
				}()

				// 更新活动时间
				GlobalMonitor.UpdateActivity()

				err := t.Download()
				if err != nil {
					Error(i18n.T("download_error", t.FullPath, err))
					_ = os.Remove(t.FullPath)
					// 更新活动时间
					GlobalMonitor.UpdateActivity()
					// 调用失败回调
					if t.OnFailure != nil {
						t.OnFailure(t.Url, t.SavePath, t.FileName, err)
					}
					return
				}
				// 更新活动时间
				GlobalMonitor.UpdateActivity()
				Success(i18n.T("download_completed", t.FullPath))
			}(t)
		}
	}()
}
