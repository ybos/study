package main

import (
	"fmt"
)

func main() {
	str1 := "abcedf"
	str2 := "abc一二三"

	fmt.Println(str1[:4])
	fmt.Println(str2[:4], str2[:6]) // 截取的是字节长度，所以中文会有问题

	_str2 := []rune(str2)
	fmt.Println(string(_str2[:4])) // 转换为rune类型再截取
}
