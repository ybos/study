package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用 goroutine 的数量
	taskLoad         = 10 // 要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建一个有缓存的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task: %d", post)
	}

	// 当所有工作都处理完成后关闭通道
	// 以便所有 goroutine 退出
	// 这里关闭通道非常重要，通道被关闭后，依然可以从通道中获取值，但是不能写入新数据
	// 这保证了通道不会因为异常关闭而丢失数据
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks

		if !ok {
			// 通道已经为空，并且已经被关闭
			fmt.Printf("Worker: %d: shutting down\n", worker)
			return
		}

		// 显示我们已经开始工作
		fmt.Printf("Worker %d: started:\t %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们已经完成了工作
		fmt.Printf("Worker %d: Completed:\t %s\n", worker, task)
	}
}
