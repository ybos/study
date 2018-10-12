package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const apple string = "苹果"
const pear = "梨"

const (
	orange string = "橘子"
	banana = "香蕉"
)

const grape, kiwi = "葡萄", "猕猴桃"

const watermelon = 100
const lemon, mango = 200, "芒果"

const durian = unsafe.Sizeof(lemon)


const pen = iota
const pencil = iota

const (
	ball_pen1 = iota
	ball_pen2 = iota
)

const (
	nephew = iota
	niece = iota
	_
	cousin = iota
)

const (
	coca_cola = iota
	pepsi = 3.14
	water = iota
)

const (
	bedroom = iota * 2
	kitchen = iota
	bathroom = iota
)

const (
	livingroom = iota * 2
	schoolroom
	fuckingroom
)

const (
	a, b = iota, iota + 5
	c, d
	e = iota
)

func main() {
	fmt.Println("apple: ", apple, "\ttype: ", reflect.TypeOf(apple))
	fmt.Println("pear: ", pear, "\ttype: ", reflect.TypeOf(pear))
	fmt.Println("orange: ", orange, "\ttype: ", reflect.TypeOf(orange))
	fmt.Println("banana: ", banana, "\ttype: ", reflect.TypeOf(banana))
	fmt.Println("grape: ", grape, "\ttype: ", reflect.TypeOf(grape))
	fmt.Println("kiwi: ", kiwi, "\ttype: ", reflect.TypeOf(kiwi))

	fmt.Println("watermelon: ", watermelon, "\ttype: ", reflect.TypeOf(watermelon))
	fmt.Println("lemon: ", lemon, "\ttype: ", reflect.TypeOf(lemon))
	fmt.Println("mango: ", mango, "\ttype: ", reflect.TypeOf(mango))

	fmt.Println("durian: ", durian, "\ttype: ", reflect.TypeOf(durian))

	fmt.Println("-------- iota --------")
	fmt.Println("pen: ", pen, "\tpencil: ", pencil)
	fmt.Println("ball pen 1: ", ball_pen1, "\tball pen 2: ", ball_pen2)
	fmt.Println("nephew: ", nephew, "\tniece: ", niece, "\tcousin: ", cousin)
	fmt.Println("coca_cola: ", coca_cola, "\tpepsi: ", pepsi, "\twater: ", water)
	fmt.Println("bedroom: ", bedroom, "\tkitchen: ", kitchen, "\tbathroom: ", bathroom)
	fmt.Println("livingroom: ", livingroom, "\tschoolroom: ", schoolroom, "\tfuckingroom: ", fuckingroom)
	fmt.Println("a: ", a, "\tb: ", b, "\tc: ", c, "\td: ", d, "\te: ", e)
}
