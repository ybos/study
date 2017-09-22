package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析参数， 默认是不会解析的
	// 如果不调用 ParseForm() 则不会向 Form 中写入映射字段
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Println("URL.Scheme: ", r.URL.Scheme)
	fmt.Println("URL.Host: ", r.URL.Host)
	fmt.Println("URL.Hostname: ", r.URL.Hostname())
	fmt.Println("URL.Port: ", r.URL.Port())
	fmt.Println("URL.Path: ", r.URL.Path)
	fmt.Println("URL.Query: ", r.URL.Query())

	fmt.Println("-----------------")

	//	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Printf("Key: %s\tValue: %s\n", k, v)
	}

	fmt.Println("-----------------")

	//	for k, v := range r.URL {
	//		fmt.Printf("Key: %s\tValue: %s\n", k, v)
	//	}

	//	fmt.Println("-----------------")

	// 将需要返回给浏览器的信息写入到响应体中
	fmt.Fprint(w, "Hello, Neil!")
}

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHelloName)

	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("Listen And Serv: ", err)
	}
}
