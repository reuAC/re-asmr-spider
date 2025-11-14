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
}

func NewWorkerPool(WorkerCount int) *WorkerPool {
	return &WorkerPool{
		cond:      sync.NewCond(&sync.Mutex{}),
		Limit:     WorkerCount,
		TaskQueue: make(WorkerChan, WorkerCount),
	}
}

func (wp *WorkerPool) Start() {
	go func() {
		for t := range wp.TaskQueue {
			wp.cond.L.Lock()
			for wp.Count >= wp.Limit {
				wp.cond.Wait()
			}
			wp.Add(1)
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
