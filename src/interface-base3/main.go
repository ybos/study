package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name   string
	age    int
	status string
	tall   float64
	weight float64
}

// fmt 包中的每个函数都实现了 String() 接口函数
// 只要类型实现了 String() 接口函数，都可以传入 fmt 系的函数内
func (h *Human) String() string {
	return "name is " + h.name + " - age is " + strconv.Itoa(h.age) + " - status is " + h.status + " - tall is " + strconv.FormatFloat(h.tall, 'f', 2, 64) + "cm - weight is " + strconv.FormatFloat(h.weight, 'f', 2, 64) + "kg"
}

func main() {
	h := Human{"Neil", 27, "alive", 178.5, 65}

	fmt.Printf("This human is ", &h)
}
