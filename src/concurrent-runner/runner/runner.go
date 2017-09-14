package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	// 用来收发 os.Signal 接口类型的值
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	complete chan error

	// timeout 报告处理任务已经超时
	// 这是一个只读的通道
	// 只读通道 channel := <-chan int
	// 只写通道 channel := chan<- int
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// 设置一个超时的错误
var ErrTimeout = errors.New("received timeout")

// 设置一个操作系统的时间返回错误
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// 语言会在指定的 duration 事件到期后,向通道发送一个 time.Time 的值
		timeout: time.After(d),
	}
}

// Add 将一个任务附加到 Runner 上,这个任务是一个接受一个 int 类型的 ID 作为参数的函数
// 这里是可变参数,接收任意多个"传入一个整型变量且无返回值的函数",这些函数都将保存在 tasks 切片中
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务,并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发送的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.goInterrupt() {
			return ErrInterrupt
		}

		// 执行已经注册的任务
		task(id)
	}

	return nil
}

// goInterrupt 验证是否接收到中断信号
func (r *Runner) goInterrupt() bool {
	select {
	// 当中断事件被出发时发出信号
	// 因为 interrupt 为空,这里在获取通道的时候,会阻塞等待
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	// 原本在接收信号的时候会阻塞,default 分支将阻塞变成了非阻塞
	// 有信号的时候会处理, 没有信号的时候会执行 default 分支
	default:
		return false
	}
}
