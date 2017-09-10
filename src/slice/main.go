package main

import (
	. "fmt"
)

func main() {
	// 指定切片的长度和容量都是5
	var slice1 = make([]int, 5)

	slice1[0] = 10
	slice1[1] = 20
	slice1[2] = 30
	slice1[3] = 40
	slice1[4] = 50

	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))

	for k := 0; k < 5; k++ {
		Println(slice1[k])
	}

	Println("\r\n")

	// 向slice1添加十个元素
	slice1 = append(slice1, 60)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 70)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 80)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 90)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 100)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 110)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 120)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 130)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 140)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 150)
	Printf("slice1 points to: %p\tlen: %d\tcap: %d\n", &slice1, len(slice1), cap(slice1))

	for k := 0; k < 15; k++ {
		Println(slice1[k])
	}

	Println("\r\n")

	// 指定不一样的长度和容量
	var slice2 = make([]int, 2, 5)

	slice2[0] = 10
	slice2[1] = 20

	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))

	// 虽然容量是5，代表了切片的底层数组可以使用5个元素，但是可见的只有2个，其他3个元素为不可见，需要调用append来改变切片大小
	for k := 0; k < 2; k++ {
		Println(slice2[k])
	}

	// 向slice2添加五个元素
	slice2 = append(slice2, 60)
	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 70)
	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 80)
	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 90)
	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 100)
	Printf("slice2 points to: %p\tlen: %d\tcap: %d\n", &slice2, len(slice2), cap(slice2))

	for k := 0; k < 7; k++ {
		Println(slice2[k])
	}

	Println("\r\n")

	// 通过初始化数据来决定切片大小
	slice3 := []int{10, 20, 30}

	for k := 0; k < 3; k++ {
		Println(slice3[k])
	}

	Println("\r\n")

	// 使用切片来赋值切片
	// 使用切片来赋值新切片，两个切片将共享底层数组，但是其他元素不可见
	slice4 := []int{10, 20, 30, 40, 50}
	slice5 := slice4[1:3]

	for k := 0; k < 5; k++ {
		Println("slice4[", k, "]: ", slice4[k])
	}

	for k := 0; k < 2; k++ {
		Println("slice5[", k, "]: ", slice5[k])
	}

	slice5[1] = 100

	Println("change slice5[1]")

	for k := 0; k < 5; k++ {
		Println("slice4[", k, "]: ", slice4[k])
	}

	for k := 0; k < 2; k++ {
		Println("slice5[", k, "]: ", slice5[k])
	}

	Println("\r\n")
	slice6 := []int{1, 2}
	slice7 := []int{3, 4}

	// 这里使用...运算符，可以将一个切片的所有元素追加到另一个切片中
	slice6 = append(slice6, slice7...)
	Println("merge slice: ", slice6)

	Println("\r\n")

	// 迭代器
	// for range 语法的使用

	for index, value := range slice6 {
		Println("slice6 index: ", index, " value: ", value)
	}
}
