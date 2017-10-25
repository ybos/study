package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func init() {
	// 初始化设置
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	// 设置路由规则
	http.HandleFunc("/page-visit", pageVisit)

	fmt.Println("start server")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
