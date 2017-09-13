package main

import (
	"fmt"
)

type notification interface {
	notify()
}

// 这里声明了一个类型 user
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("Sending email to ", u.name, "<", u.email, ">")
}

// 这里的 admin 声明为一个外部类型，内嵌一个内部类型 user
type admin struct {
	// 这里就是嵌入类型
	// 嵌入类型的所有标识符包括：方法、变量等都将提升到和外部类型一个级别
	user

	level int
}

// 这里类型既覆盖了内部类型的变量，又覆盖了内部类型的方法
type superAdmin struct {
	user

	level int

	email string
}

func (u *superAdmin) notify() {
	fmt.Println("Sending superAdmin email to ", u.name, "<", u.email, ">")
}

func main() {
	user1 := user{"John", "john@email.com"}

	// 外部类型的初始化方式一
	admin1 := admin{
		user: user{
			name:  "Neil",
			email: "neil@email.com",
		},

		level: 10,
	}

	user1.notify()

	// 既可以通过内部类型来调取方法
	admin1.user.notify()
	// 也可以直接通过调取内部类型的方法来调用，因为内部类型的所有标识符都会提升到外部类型上
	admin1.notify()

	// 提升的内容不光包括方法，还有变量等
	fmt.Println("user1.name: ", user1.name)
	fmt.Println("admin1.name: ", admin1.name)

	// 外部类型的初始化方法二
	sp := superAdmin{user{"Sofia", "sofia@email.com"}, 100, "sofia@example.com"}

	sp.notify()
}
