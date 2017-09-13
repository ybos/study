package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown 是通知正在执行的 goroutine 停止工作的标识
	shutdown int64

	// wg 用来计数
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行的时间
	time.Sleep(1 * time.Second)

	// 停止工作,安全的设置 shutdown 标识
	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 检测是否需要停止工作
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
