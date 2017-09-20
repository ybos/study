package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I'm %s, how are you?\n", h.name)
}

func (h *Human) Sing() {
	fmt.Printf("I like singing music, but I'm %d years old.\n", h.age)
}

type Man interface {
	SayHi()
	Sing()
}

func main() {
	john := Student{Human{"John", 15}, "SIIT"}
	neil := Employee{Human{"Neil", 27}, "Ecovacs"}

	// 任意实现了接口的接口值，都可以复制给接口变量
	var i Man
	i = &john

	i.SayHi()
	i.Sing()

	// 任意实现了接口的接口值，都可以复制给接口变量
	var u Man
	u = &neil

	u.SayHi()
	u.Sing()

	// 空接口 interface{} 不包含任何的 method
	// 正因如此，所以可以储存任何的值，类似 C 语言的 void * 类型指针
	// 可以定义，但是无法使用。
	//	var empty interface{}
	//	empty = neil
}
