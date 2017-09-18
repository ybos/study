package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	// New 函数的原型
	// func New(out io.Writer, prefix string, flag int) *Logger

	// 输出设备为"抛弃"的
	// ioutil.Discard 是一个 io.Writer 的实现, 所有 write 都不会有所动作, 但是会返回成功
	Trace = log.New(ioutil.Discard, "Trace:", log.Ldate|log.Ltime|log.Lshortfile)

	// 标准输出
	// Stdin/Stdout/Stderr 是已经打开的文件,分别指向标准输入,标准输出和标准错误的文件描述符
	// var (
	//     Stdin   = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	//     Stdout  = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	//     Stderr  = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	// )
	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)

	// 标准输出
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)

	// 写文件+错误输出
	// io.MultiWriter 是一个返回 io.Writer 接口类型值的可变参数函数
	// 当向其写入的时候,会向所有绑定在一起 io.Writer 同时写入
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")

	Info.Println("Special Information")

	Warning.Println("There is something you need to know about")

	Error.Println("Something has failed")
}
