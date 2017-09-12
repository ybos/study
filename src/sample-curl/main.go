package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example <url>")
		os.Exit(-1)
	}
}

func main() {
	// 从 Web 服务器获取数据
	// http 的 Get 函数返回一个 http.Response 的指针，指向获取到的信息
	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println("err")

		os.Exit(-1)
	}

	// 把数据从 Body 拷贝到标准输出 Stdout
	// Copy 函数第一个参数需要实现 io.Writer 接口
	// Copy 函数第二个参数需要实现 io.Reader 接口
	io.Copy(os.Stdout, r.Body)

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)

		os.Exit(-1)
	}
}
