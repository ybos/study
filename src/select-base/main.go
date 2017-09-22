package main

import (
	"fmt"
)

func finbonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		// select 会阻塞等待任意 case 准备就绪, 就执行准备就绪的 case
		// 如果同时多个 case 准备就绪, 随机取一个
		// defualt 语法较为特殊, 就是当监听的 case 都没准备好时执行 default
		// default 使阻塞的 select 变得不阻塞
		select {
		// 当通道 c 可写入的时候,写入一个值
		case c <- x:
			x, y = y, x+y
		// 当 quit 可读的时候,读取并退出
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		// 分 10 次读取数据
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		// 读取数据结束之后,
		// 给一个可以退出的标识
		quit <- 0
	}()

	// 执行写操作
	finbonacci(c, quit)
}
