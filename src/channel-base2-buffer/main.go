package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func read1(c chan int) {
	defer w.Done()

	// 当通道关闭后,for 循环也会被关闭
	for v := range c {
		fmt.Println("Read1 & Channel: ", v)
		time.Sleep(time.Second)
	}
}

func read2(c chan int) {
	defer w.Done()

	// 当通道关闭后,for 循环也会被关闭
	for v := range c {
		fmt.Println("Read2 & Channel: ", v)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	w.Add(2)

	// 有缓存通道最多缓存5个数据
	// 在有空间的情况下,写操作可以无阻塞操作
	// 但是在没有空间的情况下,会阻塞操作直到可以写入
	// 读操作可以无阻塞的读取数据,但是没有数据的情况下会阻塞等待读取
	c := make(chan int, 5)

	c <- 1
	fmt.Println("insert: 1")
	c <- 2
	fmt.Println("insert: 2")
	c <- 3
	fmt.Println("insert: 3")

	go read1(c)
	go read2(c)

	c <- 4
	fmt.Println("insert: 4")
	c <- 5
	fmt.Println("insert: 5")
	c <- 6
	fmt.Println("insert: 6")
	c <- 7
	fmt.Println("insert: 7")
	c <- 8
	fmt.Println("insert: 8")
	c <- 9
	fmt.Println("insert: 9")
	c <- 10
	fmt.Println("insert: 10")

	close(c)

	w.Wait()
	fmt.Println("finished")
}
