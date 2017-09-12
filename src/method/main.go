package main

import (
	"fmt"
)

// 使用指针接收者，只能传递指针
// 使用值接收者，可以传递指针或者值
// Method Receivers       | values
// -------------------------------
// (t T)                    T and *T
// (t *T)                   *T

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// 定义了一个人的结构体
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	// 创建一个 user 类型的值， 并发送通知

	u := user{"John", "John@email.com"}

	// 因为 notifier 是接口，无所谓传递的是指针还是值
	// 所以这里传递一个 user 类型的指针给函数，来实现指针接收者的限制
	sendNotification(&u)
}

// sendNotification 接收一个实现了 notifier 接口的值
// 在 Go 语言中，并不需要明确写继承关系，只需要实现一个接口所有的方法就可以多态
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}
