package main

import (
	"fmt"
)

func main() {
	var array1 [5]int

	array1[0] = 10
	array1[3] = 20
	// array1[5] = 10 // 越界编译器会报错中断

	for k := 0; k < 5; k++ {
		fmt.Println(array1[k])
	}

	fmt.Println("\r\n")

	// 声明并初始化数组
	array2 := [5]int{10, 20, 30, 40, 50}

	for k := 0; k < 5; k++ {
		fmt.Println(array2[k])
	}

	fmt.Println("\r\n")

	// 初始化部分数据
	array3 := [5]int{1: 10, 2: 20}

	for k := 0; k < 5; k++ {
		fmt.Println(array3[k])
	}

	fmt.Println("\r\n")

	// 根据初始化数据长度决定容器大小
	array4 := [...]int{1, 2, 3, 4}

	for k := 0; k < 4; k++ {
		fmt.Println(array4[k])
	}

	fmt.Println("\r\n")

	// 指针数组
	// 这里一定要创建存储元素的内存块，否则将被初始化为0值而无法使用
	array5 := [5]*int{new(int), new(int), new(int), new(int), new(int)}

	*array5[0] = 10
	*array5[1] = 20
	*array5[2] = 30
	*array5[3] = 40
	*array5[4] = 50

	for k := 0; k < 5; k++ {
		fmt.Printf("%p : %d\n", array5[k], *array5[k])
	}

	fmt.Println("\r\n")

	// 数组的复制
	var array6 [5]*int

	array6 = array5

	*array6[3] = 100

	for k := 0; k < 5; k++ {
		fmt.Printf("%p : %d\n", array6[k], *array6[k])
	}

	// 数组长度不同，类型不同将无法复制数组
	//	var array7 [4]int
	//	array7 = array1

	fmt.Println("\r\n")

	var array8 [5]int

	array8 = array1

	for k := 0; k < 5; k++ {
		fmt.Printf("array1[%d] : %p : %d\n", k, &array1[k], array1[k])
		fmt.Printf("array8[%d] : %p : %d\n", k, &array8[k], array8[k])
	}
}
