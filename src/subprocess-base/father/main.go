package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Super process:", os.Getpid())

	// 调用别的程序
	cmd := exec.Command("../child/child.exe")
	// 1. 手动调用 Start() 和 Wait()
	// 2. 调用 Run(), 函数内部将调用 Start() 和 Wait()
	// 3. 调用Output(), 内部调用 Run(), 并且将返回内容返回, 要求 stdout 必须为空, 函数会将 stdout 赋值为一个 buf
	content, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(content))
}
