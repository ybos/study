package main

import (
	"fmt"

	"permission/counter"
)

func main() {

	// NoticeCounter 因为首字母是大写的，所以是包外可见的公开的标识符
	num1 := counter.NoticeCounter(10)

	// alterCounter 因为首字母是小写的，所以是包外不可见的未公开的标识符
	//	num2 := counter.alterCounter(10)

	// 调用了一个公开的函数用来返回包内不可见的未公开变量
	// 首先这个函数也必须是公开可访问的标识符
	num3 := counter.NewAlter(100)

	fmt.Printf("notice counter: %d\n", num1)
	fmt.Printf("alert counter: %d\n", num3)

	fmt.Println("\r\n")

	// 这里只能通过默认初始化给email字段赋零值
	test1 := counter.Test{
		Name: "Neil",
	}

	// Name字段因为是公开的，所以可以使用
	fmt.Println("Name is ", test1.Name)
	// email字段因为未公开，只能使用默认初始化的零值，且外部无法访问
	//	fmt.Println("email is ", test1.email)

	fmt.Println("\r\n")

	// 这里的 Admin 是公开的
	// Rights 也是公开的
	// user 是未公开的，无法被结构字面量初始化，但是创建的时候会被默认值初始化，所以可以继续使用
	// 后续在使用的时候，因为内部类型的标识符会被提升到外部类型上，所以可以通过外部类型直接访问
	admin1 := counter.Admin{
		Rights: 10,
	}

	admin1.Name = "Neil"
	admin1.Email = "neil@email.com"

	fmt.Printf("Name: %s\tEmail: %s\tRights: %d\n", admin1.Name, admin1.Email, admin1.Rights)
}
