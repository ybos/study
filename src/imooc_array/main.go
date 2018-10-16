package main

import "fmt"

// it's copy the value to the function
func printArray(arr [5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println("print arr: ", i, v)
	}
}

// also we can use a pointer which is point to an array
func printAnyArray(arr *[5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println("print arr: ", i, v)
	}
}

func main() {
	var arr1 [5]int

	var arr2 = [5]int {1, 2, 3, 4, 5}

	var arr3 = [...]int {2, 3, 4}

	var arr4 = [3]int {1: 11, 0: 22}

	fmt.Println(arr1, arr2, arr3, arr4)

	for i := 0; i < 3; i++ {
		fmt.Println("loop arr4: ", i, arr4[i])
	}

	fmt.Println("-----------------")

	for i := range arr3 {
		fmt.Println("loop arr3: ", i, arr3[i])
	}

	fmt.Println("-----------------")

	for i, v := range arr2 {
		fmt.Println("loop arr2: ", i, v)
	}

	fmt.Println("-----------------")

	for i := 0; i < len(arr1); i++ {
		fmt.Println("loop arr1: ", i, arr1[i])
	}

	fmt.Println("-----------------")

	printArray(arr2)

	fmt.Println("-----------------")    

	for i, v := range arr2 {
		fmt.Println("loop arr2: ", i, v)
	}

	fmt.Println("-----------------")

	printAnyArray(&arr2)

	fmt.Println("-----------------")

	for i, v := range arr2 {
		fmt.Println("loop arr2: ", i, v)
	}
}
