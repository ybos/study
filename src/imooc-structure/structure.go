package main

import (
	"fmt"
)

// we usually use capital words to define constant
const WEBSITE string = "https://www.imooc.com"
const TITLE string = "慕课网"

// we usually use lower case to define variable
var visitCounter int = 0

// this means customType is an alias for string
// both them are similar, but customType and string are different type in Golang.
type customType string

// structure define
type Learn struct {
	id int
	name string
}

// interface define
type ILearn interface {

}

func learnGo () {
	fmt.Println("Hello, Golang!")
}

func main() {
	fmt.Println("Hello, main!")

	learnGo()
}
