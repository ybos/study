package main

import (
	"fmt"
	"time"
)

func main() {
	var index int8 = 0
	for {
		fmt.Println(index, "iMooc")

		index++

		if index == 10 {
			break
		}
	}

	fmt.Println("=================")

	for index = 0; index < 10; index++ {
		fmt.Println(index, "iMooc 2")
	}

	arr := []string{"apple", "pear", "orange", "banana"}

	for key, value := range arr {
		fmt.Println(key, " : ", value)
	}

	goto One

	fmt.Println("Start test")

	One:
		fmt.Println("Print code block")

	index  = 0

	for {
		fmt.Println("<code> Hello, World </code>")

		time.Sleep(time.Second)

		index++

		if index == 5 {
			goto Two
		}
	}

	Two:

	fmt.Println("ends")

	for index = 0; index <= 10; index++{
		if index % 2 == 0 {
			fmt.Println("it's an even number")
			continue;
		}

		fmt.Println("number is ", index)
	}
}
