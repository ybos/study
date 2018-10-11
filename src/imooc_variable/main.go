package main

import (
	"fmt"
	"reflect"
)

// define a global variable, must use var keyword
var s string
// you can use the variable which the first letter is capitial in other package.
var Str string

func main() {
	// define a variable
	var i8 int8

	// set a value
	i8 = 32

	// define and set a variable
	var i32 int32 = 32

	fmt.Println(i8, i32, s, Str)

	// define some variables in a block
	var (
		f32 float32 = 123.12
		f64 float64 = 321.32
	)

	fmt.Println(f32, f64)

	// simple define and set
	shortVariable := 123

	fmt.Println("shortVariable's value: ", shortVariable, "\tshortVariable's type: ", reflect.TypeOf(shortVariable), "\n")

	// convert
	i32 = int32(f32)
	fmt.Println(i32)
}
