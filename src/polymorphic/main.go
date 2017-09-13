package main

import (
	"fmt"
)

/**
 * 这里只定义了接口的相关行为
 * 只要满足接口的全部行为，就可以认为是某个类型继承自本接口
 */
type notification interface {
	notify()
}

/**
 * 声明独立的类型 user
 */
type user struct {
	name  string
	email string
}

/**
 * 类型 user 有一个指针接收者的方法
 */
func (u *user) notify() {
	fmt.Printf("Sending email to user %s<%s>\n", u.name, u.email)
}

/**
 * 声明了独立的类型 admin
 */
type admin struct {
	name  string
	email string
}

/**
 * 类型 admin 有一个指针接收者的方法
 */
func (a *admin) notify() {
	fmt.Printf("Sending email to administrator %s<%s>\n", a.name, a.email)
}

func main() {
	user1 := user{"John", "john@email.com"}
	admin1 := admin{"Neil", "neil@email.com"}

	// 两个不同的类型传入同一个函数
	sendEmail(&user1)
	sendEmail(&admin1)

}

/**
 * 函数的参数是一个接口
 * 意思是只要满足本接口的行为定义，任何类型的值都可以传入，包括指针和值的区别都可以被忽略
 */
func sendEmail(u notification) {
	// 这里是通过接口调用的函数，无关具体类型
	u.notify()

	// 这里调用成员属性无效，报错为未定义：u.name undefined (type notification has no field or method name)
	//	fmt.Println("test if I can receive member var: ", u.name)
}
