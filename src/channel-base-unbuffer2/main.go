package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	time.Sleep(5 * time.Second)

	// 使用无缓存通道有个特色就是必须写入通道和读取通道同时准备好才会生效
	// 既上面 goroutine 已经准备好5秒都未曾输出东西,是因为下面的写入通道未准备就绪
	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	// 开始围绕跑道跑步
	fmt.Printf("Runner %d Running with baton\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		// 这里即便开启了一个新的 goroutine 但是得到下面写入通道准备就绪后才可正式开始执行， 读取通道一开始就阻塞了
		go Runner(baton)
	}

	time.Sleep(1000 * time.Millisecond)

	// 判断比赛结束
	if runner == 4 {
		fmt.Printf("Runner %d finished, race over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)

	baton <- newRunner
}
