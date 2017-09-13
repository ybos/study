package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg 用来等待程序完成
var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	// 分配所有可用的物理处理器给调度器使用
	//	runtime.GOMAXPROCS(runtime.NumCPU())

	// 技术信号量为 2, 等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	// 等待 goroutine 结束
	fmt.Println("Waiting to finish")

	// 这是一个阻塞的函数,用来等待goroutine的执行完成
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// printPrime 显示5000以内的像素值
func printPrime(prefix string) {
	// 在程序退出时调用 Done 来通知 main 函数工作已经结束
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Conpleted", prefix)
}
