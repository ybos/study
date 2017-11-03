package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func mixer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.PostForm)
}

func main() {
	http.HandleFunc("/", mixer)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}

	res, err := http.PostForm("http://account.ecovacs.cn/account/member/signin", url.Values{"account": {"test"}, "password": {"test"}})

	if err != nil {
		fmt.Println("请求错误" + err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("解析错误" + err.Error())
		return
	}

	fmt.Println(string(body))
}
