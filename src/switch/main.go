package main

import "fmt"

func main() {
	var str = "test2"

	switch str {
	case "test1":
		fmt.Println("test1")
	case "test2":
		fmt.Println("test2")
	case "test3", "test4":
		fmt.Println("test3 or test4")
	default:
		fmt.Println("default value")
	}
}
