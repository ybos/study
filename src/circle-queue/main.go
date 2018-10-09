package main

import (
	"fmt"
	"circle-queue/queue"
	"time"
	"sync"
)

func main() {
	fmt.Println("test")

	d := new(queue.CircleQueue)

	//ok1, val1 := d.Get()
	//if ok1 {
	//	fmt.Println("获取的值是：", val1)
	//} else {
	//	fmt.Println("获取失败")
	//}
	//
	//d.Put("first")
	//d.Put("second")
	//d.Put("third")
	//
	//ok2, val2 := d.Get()
	//if ok2 {
	//	fmt.Println("获取的值是：", val2)
	//} else {
	//	fmt.Println("获取失败")
	//}
	//
	//d.Put("fourth")
	//d.Put("fifth")
	//d.Put("sixth")
	//d.Put("seventh")
	//d.Put("eighth")
	//d.Put("ninth")
	//d.Put("tenth")
	//d.Put("eleventh")
	//
	//ok3, val3 := d.Put("twelfth")
	//if ok3 {
	//	fmt.Println("存储成功，队列内有：", val3)
	//} else {
	//	fmt.Println("存储失败")
	//}
	//
	//ok4, val4 := d.Get()
	//if ok4 {
	//	fmt.Println("获取的值是：", val4)
	//} else {
	//	fmt.Println("获取失败")
	//}
	//
	//ok5, val5 := d.Put(12)
	//if ok5 {
	//	fmt.Println("存储成功，队列内有：", val5)
	//} else {
	//	fmt.Println("存储失败")
	//}

	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func(d *queue.CircleQueue) {
	//	defer wg.Done()
	//
	//	start := time.Now()
	//
	//	for i := 0; i < 1000000; i++ {
	//		d.Put(i)
	//	}
	//
	//	end := time.Now()
	//	use := end.Sub(start)
	//
	//	fmt.Println("Put method, use ", use)
	//}(d)
	//
	//go func(d *queue.CircleQueue) {
	//	defer wg.Done()
	//
	//	start := time.Now()
	//
	//	for i := 0; i < 1000000; i++ {
	//		d.Get()
	//	}
	//
	//	end := time.Now()
	//	use := end.Sub(start)
	//
	//	fmt.Println("Get method, use ", use)
	//}(d)
	//
	//wg.Wait()

	/* 检查在无锁状态下，是否会有数据丢失 */
	var wg sync.WaitGroup
	wg.Add(3)

	c1, c2, c3 := 0, 0, 0

	start := time.Now()
	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			d.Put(1)
			c1++
		}
	}()

	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			d.Put(2)
			c2++
		}
	}()

	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			d.Put(3)
			c3++
		}
	}()

	wg.Wait()
	end := time.Now()
	fmt.Println("ninety thousand times spends: ", end.Sub(start))

	total, i1, i2, i3 := 0, 0, 0, 0

	for true {
		ok, val := d.Get()

		total++

		if ok {
			if val == 1 {
				i1++
			} else if val == 2 {
				i2++
			} else if val == 3 {
				i3++
			}
		} else {
			break
		}
	}

	fmt.Println("loop c1: ", c1, "\tc2: ", c2, "\tc3: ", c3)
	fmt.Println("total: ", total, "\ti1: ", i1, "\ti2: ", i2, "\ti3: ", i3)

	/* chan 做对比 */
	testChan := make(chan int, 100000)

	wg.Add(3)

	c1, c2, c3 = 0, 0, 0

	start = time.Now()
	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			testChan <- 1
			c1++
		}
	}()

	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			testChan <- 2
			c2++
		}
	}()

	go func () {
		defer wg.Done()

		for i := 0; i < 30000; i++ {
			testChan <- 3
			c3++
		}
	}()

	wg.Wait()
	end = time.Now()
	fmt.Println("ninety thousand times use channel spends: ", end.Sub(start))
}
