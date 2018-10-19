package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var s string = "Yes,我爱我自己！" // UTF-8 every byte need 3 bytes.

	fmt.Println(s, "\nlen(s): ", len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println("\n-------")

	for i, ch := range s { // ch is rune (an alias for int32)
		fmt.Printf("(%d, %X)", i, ch) // unicode
	}

	fmt.Println("\n-------")

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c %d)", i, ch, unsafe.Sizeof(ch))
	}
}
