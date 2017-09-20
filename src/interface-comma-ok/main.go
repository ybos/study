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
	a := Animal{"Monkey"}

	var action Action

	h.SayHi()
	a.SayHi()

	action = h

	// 使用 element.(T) 的方式进行断言式判断接口变量类型
	if value, ok := action.(Animal); ok {
		fmt.Println("h is animal, h's name is ", value.name)
	} else {
		fmt.Println("h doesn't a animal")
	}

	// 使用 element.(T) 的方式进行断言式判断接口变量类型
	if value, ok := action.(Human); ok {
		fmt.Println("h is animal, h's name is ", value.name)
	} else {
		fmt.Println("h doesn't a animal")
	}
}
