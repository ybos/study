package main

import (
	"fmt"
	"circle-queue/queue"
)

func main() {
	fmt.Println("test")

	d := new(queue.CircleQueue)

	ok1, val1 := d.Get()
	if ok1 {
		fmt.Println("获取的值是：", val1)
	} else {
		fmt.Println("获取失败")
	}

	d.Put("first")
	d.Put("second")
	d.Put("third")

	ok2, val2 := d.Get()
	if ok2 {
		fmt.Println("获取的值是：", val2)
	} else {
		fmt.Println("获取失败")
	}

	d.Put("fourth")
	d.Put("fifth")
	d.Put("sixth")
	d.Put("seventh")
	d.Put("eighth")
	d.Put("ninth")
	d.Put("tenth")
	d.Put("eleventh")

	ok3, val3 := d.Put("twelfth")
	if ok3 {
		fmt.Println("存储成功，队列内有：", val3)
	} else {
		fmt.Println("存储失败")
	}

	ok4, val4 := d.Get()
	if ok4 {
		fmt.Println("获取的值是：", val4)
	} else {
		fmt.Println("获取失败")
	}

	ok5, val5 := d.Put(12)
	if ok5 {
		fmt.Println("存储成功，队列内有：", val5)
	} else {
		fmt.Println("存储失败")
	}
}
