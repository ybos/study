package work

import "sync"

// Worker 必须满足的接口
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// 创建一个新的 Goroutine 池
func New(maxGoroutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutine)

	for i := 0; i < maxGoroutine; i++ {
		go func() {
			// 循环阻塞,直到从通道内接受到一个值,并执行相关的 Task 方法
			// 如果通道关闭,则 for 循环会结束, 并调用 Done 方法
			for w := range p.work {
				w.Task()
			}

			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作到 Goroutine 池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 Goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
