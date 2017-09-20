package main

import "fmt"

type Human struct {
	name string
}

type Animal struct {
	name string
}

func (h Human) SayHi() {
	fmt.Println("Hello, My name is ", h.name)
}

func (a Animal) SayHi() {
	fmt.Println("Hello, My name is ", a.name)
}

type Action interface {
	SayHi()
}

func main() {
	h := Human{"Neil"}
	a := Animal{"Cat"}

	var action Action

	action = h

	// 使用 switch & comma 的方法，type 只能存在于 switch 语法中
	// 如果要在外部使用，使用普通的 comma-ok 方式即可
	switch value := action.(type) {
	case Human:
		fmt.Println("This is Human ", value.name)
		break
	case Animal:
		fmt.Println("This is Animal ", value.name)
		break
	}

	action = a

	switch value := action.(type) {
	case Human:
		fmt.Println("This is Human ", value.name)
		break
	case Animal:
		fmt.Println("This is Animal ", value.name)
		break
	}
}
