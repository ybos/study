package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 设置了最多使用一个逻辑处理器来处理 goroutine
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成的计数器
	// WaitGroup 是一个计数信号量,如果 WaitGroup 的值 >0,Wait 方法就会阻塞
	var wg sync.WaitGroup
	// 计数加2,标识要等待两个 goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明了一个匿名函数,并创建一个goroutine
	go func() {
		// 在函数退出调用时,调用 Done 来通知 main 函数工作已经完成
		// 并将 wg 计数器减一
		// defer 声明在函数退出时调用 Done 方法
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println("")
		}
	}()

	// 声明了一个匿名函数,并创建一个goroutine
	go func() {
		// 在函数退出调用时,调用 Done 来通知 main 函数工作已经完成
		// 并将 wg 计数器减一
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println("")
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting to finish")

	// 这是一个阻塞的函数,用来等待goroutine的执行完成
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
