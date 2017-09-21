package main

import "fmt"

type Interface1 interface {
	SayAge()
}

// Interface2 内嵌了 Interface1
// Interface2 隐士的包含了 Interface1 的全部接口
type Interface2 interface {
	Interface1
	SayHello()
}

type Age int
type Name string

func (i Age) SayAge() {
	fmt.Printf("Hi, I'm %d years old.\n", i)
}

func (n Name) SayAge() {
	fmt.Printf("Hello, I'm %s years old.\n", n)
}

func (n Name) SayHello() {
	fmt.Printf("Hello, My name is %s.\n", n)
}

func main() {
	age := Age(27)
	name := Name("Neil")

	age.SayAge()
	name.SayAge()
	name.SayHello()

	var i1 Interface1

	i1 = name
	i1.SayAge()

	var i2 Interface2
	// 不能用 age 赋值给 Interface2，因为 age 没有实现 Interface2 的全部接口
	//	i2 = age
	i2 = name
	i2.SayHello()
}
