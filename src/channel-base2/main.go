package main

import (
	. "fmt"
)

func sum(a []int, c chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	for _, v := range a[:len(a)/2] {
		Println(v)
	}

	Println("----")

	for _, v := range a[len(a)/2:] {
		Println(v)
	}

	c := make(chan int)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)

	// 阻塞的接收两次通道内的值
	x, y := <-c, <-c

	Println("----")

	Println(x, y, x+y)
}
