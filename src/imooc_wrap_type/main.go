package main

import "fmt"

type Queue []int
type Queue2 []int

// Cannot change the original data
func (q Queue2) Push(v int) {
	q = append(q, v)
}

// 压入值
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 从第一个弹出值
func (q *Queue) Pop() int {
	head := (*q)[0]

	*q = (*q)[1:]

	return head
}

// 判断是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func main() {
	q := Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println("--------------")

	q2 := Queue2{1}

	q2.Push(2)
	q2.Push(3)

	fmt.Println(q2)
}
