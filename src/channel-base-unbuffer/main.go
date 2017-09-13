package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 初始化函数,在main之前执行
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 初始化一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	// 初始化两个
	go player("John", court)
	go player("Neil", court)

	// 发球
	// 必须有一个起始的数据,否则两个goroutine都在等待,则会终止
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court

		// 如果通道被关闭,则算获胜
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数,然后判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 并关闭通道
			close(court)
			return
		}

		// 显示击球数,并将击球数 +1
		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++

		// 将球打向对手
		court <- ball
	}
}
