package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int

	// 计数器
	wg sync.WaitGroup

	// 互斥锁,用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
}

// incCounter 使用互斥锁来同步保证安全访问
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个 goroutine 获得锁
		// 其他 goroutine 都将等待直到获取到锁为止
		mutex.Lock()

		// 获取外部变量
		value := counter

		// 强行从正在执行的线程中退出并加入等待队列
		runtime.Gosched()

		value++

		counter = value

		// 释放锁
		mutex.Unlock()
	}
}
