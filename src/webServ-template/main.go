package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func loginAction(w http.ResponseWriter, r *http.Request) {
	// 解析一下请求到 Form 字段内
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key: ", k, "\tvalue: ", v)

		if k == "username" {
			for sk, sv := range r.Form["username"] {
				fmt.Println("username index: ", sk, "\tvalue: ", sv)
			}
		}

		if k == "password" {
			for sk, sv := range r.Form["password"] {
				fmt.Println("password index: ", sk, "\tvalue: ", sv)
			}
		}
	}

	// 将输出内容赋值给 w
	fmt.Fprintf(w, "Hello, Neil!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)

	// 根据请求类型决定哪种处理方式
	if r.Method == "GET" {
		// 生成 md5 值
		curTime := time.Now().Unix()
		md5Value := md5.New()
		io.WriteString(md5Value, strconv.FormatInt(curTime, 10))
		token := fmt.Sprintf("%x", md5Value.Sum(nil))

		// 如果是 GET 请求，直接输出模板
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		loginAction(w, r)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {

}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request.username: ", r.FormValue("username"))
}

func main() {
	// 监控 /login 这个 path
	http.HandleFunc("/login", login)

	// 上传功能
	http.HandleFunc("/upload", upload)

	// 默认方式
	http.HandleFunc("/", defaultPage)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
