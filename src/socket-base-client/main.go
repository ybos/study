package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port \n", os.Args[0])
		os.Exit(1)
	}

	// 转换IP地址到机器地址,获取一个 tcpAddr 地址
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	// DialTCP 用于创建一个 TCP 链接的 conn
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkErr(err)

	// 向管道内发送消息
	//	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//	checkErr(err)

	// 从管道内读取消息
	result, err := ioutil.ReadAll(conn)
	checkErr(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
