//package main

//import (
//	"fmt"
//	"log"
//	"net/http"
//	"net/url"
//	"os"
//	"sync"
//	"time"

//	"analytic/sarama"
//	"analytic/user_agent"
//)

//var (
//	wg     sync.WaitGroup
//	logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)
//)

//func tongji(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("request path:\t", r.URL.Path)
//	fmt.Println("request time:\t", time.Now().Local())
//	fmt.Println("client ip:\t", r.RemoteAddr)
//	fmt.Println("param sessionid:\t", r.FormValue("sessionid"))
//	fmt.Println("param userid:\t", r.FormValue("userid"))
//	fmt.Println("param url:\t", r.FormValue("url"))
//	fmt.Println("param referer:\t", r.FormValue("referer"))

//	fmt.Println("------------")

//	u, err := url.Parse(r.FormValue("referer"))
//	if err != nil {
//		panic(err)
//	}

//	fmt.Println("param referer Host:\t", u.Hostname())

//	fmt.Println("------------")

//	u, err = url.Parse(r.FormValue("url"))
//	if err != nil {
//		panic(err)
//	}

//	fmt.Println("param url Path:\t", u.Path)

//	fmt.Println("------------")

//	ua := user_agent.New(r.UserAgent())

//	fmt.Printf("operating platform:\t%v\n", ua.Platform())
//	fmt.Printf("operating system:\t%v\n", ua.OS())

//	name, version := ua.Browser()
//	fmt.Printf("browser name:\t%v\n", name)
//	fmt.Printf("browser version:\t%v\n", version)

//	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
//}

//func main() {
//	http.HandleFunc("/visit", tongji)

//	err := http.ListenAndServe(":8080", nil)

//	if err != nil {
//		fmt.Println("Fatal error: ", err.Error())
//	}
//}

package main

import (
	"analytic/data"
	"fmt"
	"net/http"
)

//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  481.5µs
//use time:  1.5016ms
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  499.5µs
//use time:  0s
//use time:  501.1µs
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  485.5µs
//use time:  0s

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	var result = data.GetResource(r)

	//	fmt.Printf(result)

	//	fmt.Println("---------------------")

	fmt.Fprintf(w, result)
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
