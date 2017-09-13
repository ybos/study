package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int

	// wg 用来计数等待程序结束
	wg sync.WaitGroup
)

func main() {
	// 计数加2
	wg.Add(2)

	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待结束
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter(id int) {
	// 函数退出之前执行本命令
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获 counter 的值
		value := counter

		// 当前 goroutine 从线程退出,并放回队列
		// 强行将本任务从正在执行的线程退出并放回队列
		runtime.Gosched()

		// 增加本地 value 的值
		value++

		// 将该值保存回 counter
		counter = value
	}
}
