package main

import (
	"bytes"
	"fmt"
	"time"
)

// bytes.Buffer 的方式耗时几乎为0, 十万个字符的耗时为1ms
// string 的方式耗时为10ms左右

func main() {
	var b bytes.Buffer
	_b := []byte{'a'}

	t1 := time.Now()
	for i := 0; i < 10000; i++ {
		b.Write(_b)
	}
	fmt.Println("bytes: ", b.String(), "\nbytes use:", time.Since(t1))

	var s string

	t2 := time.Now()
	for i := 0; i < 10000; i++ {
		s = s + "a"
	}
	fmt.Println("string: ", s, "\nstring use:", time.Since(t2))
}
