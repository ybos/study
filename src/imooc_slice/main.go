package main

import (
	"fmt"
	"reflect"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	var arr = [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println("Base array is: ", arr)

	// [x, y)
	s := arr[2:6]

	fmt.Println("s is a ", reflect.TypeOf(s), ", and value is: ", s)

	fmt.Println("arr[2:6]:", arr[2:6])
	fmt.Println("arr[:6]:", arr[:6])
	fmt.Println("arr[2:]:", arr[2:])
	fmt.Println("arr[:]:", arr[:])

	fmt.Println("--------------")

	s1 := arr[:6]
	s2 := arr[2:6]

	fmt.Println("before s1: ", s1)
	fmt.Println("before s2: ", s2)

	updateSlice(s1)
	updateSlice(s2)

	fmt.Println("after s1: ", s1)
	fmt.Println("after s2: ", s2)
	fmt.Println("original array: ", arr)

	ss1 := s1[:3]
	sss1 := ss1[1:]

	fmt.Println("Reslice: ")
	fmt.Println(ss1, sss1)

	fmt.Println("--------------")

	var second_array = [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// second_s1 is 3,4,5,6
	second_s1 := second_array[2:6]
	second_s2 := second_s1[3:5]
	second_s3 := second_s2[1:4]

	// ptr, len, cap
	fmt.Println("second_array is ", second_array)

	fmt.Printf("second_s1 = %v, len(second_s1) = %d, cap(second_s1) = %d\n", second_s1, len(second_s1), cap(second_s1))
	fmt.Printf("second_s2 = %v, len(second_s2) = %d, cap(second_s2) = %d\n", second_s2, len(second_s2), cap(second_s2))
	fmt.Printf("second_s3 = %v, len(second_s3) = %d, cap(second_s3) = %d\n", second_s3, len(second_s3), cap(second_s3))

	fmt.Println("--------------")

	// s4/s5/s6 no longer view second_array
	// system will return a new slice include new ptr, new len, new capacity
	second_s4 := append(second_s3, 10)
	second_s5 := append(second_s4, 11)
	second_s6 := append(second_s5, 12)

	fmt.Printf("second_s4 = %v, len(second_s4) = %d, cap(second_s4) = %d\n", second_s4, len(second_s4), cap(second_s4))
	fmt.Printf("second_s5 = %v, len(second_s5) = %d, cap(second_s5) = %d\n", second_s5, len(second_s5), cap(second_s5))
	fmt.Printf("second_s6 = %v, len(second_s6) = %d, cap(second_s6) = %d\n", second_s6, len(second_s6), cap(second_s6))

	fmt.Println(second_array)

	fmt.Println("--------------")

	var third_slice []int // zero value for slice is nil

	for i := 0; i < 10; i++ {
		printSliceInfo(third_slice)

		third_slice = append(third_slice, 2 * i + 1)
	}

	fmt.Println(third_slice)

	fmt.Println("--------------")

	third_slice1 := []int {2, 4, 6, 8}
	printSliceInfo(third_slice1)

	third_slice2 := make([]int, 10, 32)
	printSliceInfo(third_slice2)

	fmt.Println("--------------")

	fourth_slice1 := []int {1, 2, 3, 4}
	var fourth_slice2 = []int {0, 0, 0, 0, 0, 0, 0, 0}
	copy(fourth_slice2, fourth_slice1)

	printSliceInfo(fourth_slice1)
	printSliceInfo(fourth_slice2)

	fmt.Println("Deleting elements from slice")
	fourth_slice2 = append(fourth_slice2[:3], fourth_slice2[4:]...)

	printSliceInfo(fourth_slice2)

	fmt.Println("Popping from front and back")
	printSliceInfo(fourth_slice2)
	fourth_front := fourth_slice2[0]
	fourth_slice2 = fourth_slice2[1:]

	fourth_tail := fourth_slice2[len(fourth_slice2) - 1]
	fourth_slice2 = fourth_slice2[:len(fourth_slice2) - 1]

	fmt.Println("front element: ", fourth_front, "\ttail element: ", fourth_tail)
	printSliceInfo(fourth_slice2)
}

func printSliceInfo(s []int) {
	fmt.Print(s, "\t")
	fmt.Printf("len = %d, capacity = %d\n", len(s), cap(s))
}