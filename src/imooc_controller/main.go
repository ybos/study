package main

import "fmt"

func main() {
	a := 1

	if a > 0 {
		fmt.Println("a is greater than 0")
	} else {
		fmt.Println("a is not greater than 0")
	}

	b := 1000

	if b > 10000 {
		fmt.Println("b > 10,000")
	} else if b >= 1000 {
		fmt.Println("b >= 1,000")
	} else if b >= 100 {
		fmt.Println("b >= 100")
	} else {
		fmt.Println("b is ", b)
	}

	var c interface{}

	c = 2

	switch c.(type) {
	case int:
		fmt.Println("case 1")
	case float32:
		fmt.Println("case 2")
	case int32:
		fmt.Println("case 3")
	default:
		fmt.Println("default value")
	}
}
