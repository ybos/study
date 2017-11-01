package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	//	"github.com/garyburd/redigo/redis"
)

func route(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Location", "/index.html")
		w.WriteHeader(http.StatusMovedPermanently)
	}

	fmt.Fprint(w, r.URL.Path)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, nil)
}

func safe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, "您请求的方式发生了错误")
	}

	urlData, err := url.Parse(r.FormValue("link"))
	if err != nil {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, "无法解析您所提交的连接")
	}

	fmt.Fprintf(w, "Hostname: %s<br />Path: %s<br />Query: %s", urlData.Hostname(), urlData.Path, urlData.Query())
}

func main() {
	http.HandleFunc("/", route)
	http.HandleFunc("/index.html", index)
	http.HandleFunc("/safe.html", safe)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
