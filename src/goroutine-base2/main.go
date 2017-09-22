package main

import (
	"fmt"
	"runtime"
)

func SayHi(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go SayHi("world")

	SayHi("hello")
}
