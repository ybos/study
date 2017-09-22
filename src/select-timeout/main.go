package main

import (
	. "fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		select {
		case v := <-c:
			Println(v)
			// time.After 会返回一个 time.C, 并且在时间到期后向管道内写入一个 time.Now() 的值
		case <-time.After(5 * time.Second):
			Println("timeout")
			o <- true
			break
		}
	}()

	// 这里只是阻塞终止而已
	<-o
}
