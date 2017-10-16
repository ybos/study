package main

import (
	"analytic-copy/data"
	"analytic-copy/db"
	"fmt"
	"net/http"
)

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	var result = data.GetResource(r)

	db.ProducerOne(result)

	//	fmt.Printf(result)

	//	fmt.Println("---------------------")

	fmt.Fprintf(w, "ok")
}

func main() {
	// 设置路由规则
	http.HandleFunc("/page-visit", pageVisit)

	fmt.Println("start server")

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
