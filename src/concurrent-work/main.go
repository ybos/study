package main

import (
	"log"
	"sync"
	"time"

	"concurrent-work/work"
)

// names 提供了一组用来显示的名字
var names = []string{
	"Steve",
	"Bob",
	"Mary",
	"Therese",
	"Jason",
}

type namePrinter struct {
	name string
}

// Task 实现了 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个 Goroutine 来创建工作池
	// 在创建工作池的时候, 工作池内的2个 Goroutine 已经开始阻塞等待工作进入了
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		// 迭代 names 切片
		for _, name := range names {
			// 创建一个 namePrinter 并提供指定的名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行, 当 Run 返回时我们就知道任务已经处理完成
				// 任务是一边提交, 一遍执行的.
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	// 等待所有提交工作的执行完成
	wg.Wait()

	// 让所有工作池停止工作, 等待现有的工作完成
	p.Shutdown()
}
