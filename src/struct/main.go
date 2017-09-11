package main

import (
	. "fmt"
)

// 创建了一个结构并命名为 user
type user struct {
	name string
	age  int
}

func main() {
	// 声明一个空的结构
	var user1 user

	// 声明并初始化一个结构
	user2 := user{
		name: "John",
		age:  15,
	}

	// 每一行的成员属性末尾都需要加逗号分隔，包括最后一行
	user3 := user{
		name: "Neil",
		age:  27,
	}

	// 通过顺序指定结构的初始化值
	user4 := user{"Sofia", 20}

	Println("user1.name: ", user1.name, "\t user1.age: ", user1.age)
	Println("user2.name: ", user2.name, "\t user2.age: ", user2.age)
	Println("user3.name: ", user3.name, "\t user3.age: ", user3.age)
	Println("user4.name: ", user4.name, "\t user4.age: ", user4.age)

	Println("\r\n")

	// 使用值接收者来调用方法
	user4.notify()

	// 使用指针接收者来
	// 使用指针来调用值接收者的方法
	user5 := &user{"Aanya", 27}
	user5.notify()

	// 不管使用值还是指针都可以实际的改变对象
	// Go语言在设计时就将两者做了转换，以极大的方便使用者
	user5.changeAge(50)
	user5.notify()
	user4.changeAge(60)
	user4.notify()
}

// 定义一个使用值接收者的方法
func (u user) notify() {
	Println("user.name: ", u.name, "\tuser.age: ", u.age)
}

// 定义一个使用指针接收者的方法
func (u *user) changeAge(newAge int) {
	u.age = newAge
}
