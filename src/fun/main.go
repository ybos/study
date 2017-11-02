package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	//	"github.com/garyburd/redigo/redis"
)

const TopDomain string = "localhost:8080"
const TopDomainLen int = 14

func route(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Location", "/index.html")
		w.WriteHeader(http.StatusMovedPermanently)
	}

	domain := []rune(r.Host)
	domainLen := len(domain)

	domain = domain[:domainLen-TopDomainLen-1]

	res, err := http.Get("http://" + string(domain) + r.URL.RequestURI())

	if err != nil {
		fmt.Fprintf(w, "请求错误(http://"+string(domain)+r.URL.RequestURI()+")"+err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(w, "解析错误(http://"+string(domain)+r.URL.RequestURI()+")"+err.Error())
		return
	}

	str := strings.Replace(string(body), string(domain), r.Host, -1)

	fmt.Fprintf(w, str)

	// 处理请求头
	//	w.Header().Set("Content-Type", res.Header["Content-Type"])
	w.Header().Set("Content-Type", res.Header["Content-Type"][0])
	fmt.Println(res.Header["Content-Type"][0])
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

	w.Header().Set("Location", "http://"+urlData.Hostname()+"."+TopDomain+urlData.RequestURI())
	w.WriteHeader(http.StatusMovedPermanently)
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
