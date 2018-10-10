package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// an alias for type
type 慕课网 int32

func main() {
	var i8 uint8 = 1
	var i32 uint32 = 1
	fmt.Println("uint8: ", unsafe.Sizeof(i8), "Byte,\tuint32: ", unsafe.Sizeof(i32), "Byte\n")

	var i int = 1
	fmt.Println("int: ", unsafe.Sizeof(i), "Byte, it depends on your system.\n")

	var b bool = true
	fmt.Println("bool: ", b, "\n")

	var c byte = 'c'
	var r rune = 'r'
	fmt.Println("byte: ", unsafe.Sizeof(c), "Byte,\trune: ", unsafe.Sizeof(r), "Byte")
	fmt.Println("byte is an alias for uint8, rune is an alias for int32.\n")

	var zero_i int32
	var zero_f float32
	var zero_b bool
	var zero_d complex64
	var zero_s string
	fmt.Println("check default value:\nint32: ", zero_i, "\tfloat32: ", zero_f, "\tbool: ", zero_b, "\tcomplex: ", zero_d, "\tstring: ", zero_s, "\n\n")

	var type_test 慕课网 = 1
	fmt.Println("alias for type: ", type_test)
	fmt.Println("type_test's type is: ", reflect.TypeOf(type_test))
	fmt.Println("慕课网 needs: ", unsafe.Sizeof(type_test), "\n")
	fmt.Println("慕课网 is an alias for int32, they are similar, but different")
}
