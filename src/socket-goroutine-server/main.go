package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"

	// 将 ip 地址转换成 tcpAddr 类型
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 监听tcp端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// 接收一个链接请求
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 启动一个 goroutine 用来处理请求
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	daytime := time.Now().String()

	// 写数据
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
