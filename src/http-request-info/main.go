package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"http-request-info/user_agent"
)

func tongji(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request path:\t", r.URL.Path)
	fmt.Println("request time:\t", time.Now().Local())
	fmt.Println("client ip:\t", r.RemoteAddr)
	fmt.Println("param sessionid:\t", r.FormValue("sessionid"))
	fmt.Println("param userid:\t", r.FormValue("userid"))
	fmt.Println("param url:\t", r.FormValue("url"))
	fmt.Println("param referer:\t", r.FormValue("referer"))

	fmt.Println("------------")

	u, err := url.Parse(r.FormValue("referer"))
	if err != nil {
		panic(err)
	}

	fmt.Println("param referer Host:\t", u.Hostname())

	fmt.Println("------------")

	u, err = url.Parse(r.FormValue("url"))
	if err != nil {
		panic(err)
	}

	fmt.Println("param url Path:\t", u.Path)

	fmt.Println("------------")

	ua := user_agent.New(r.UserAgent())

	fmt.Printf("operating platform:\t%v\n", ua.Platform())
	fmt.Printf("operating system:\t%v\n", ua.OS())

	name, version := ua.Browser()
	fmt.Printf("browser name:\t%v\n", name)
	fmt.Printf("browser version:\t%v\n", version)

	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request time:\t", time.Now().Local())
	fmt.Println("client ip:\t", r.RemoteAddr)
	fmt.Println("user-agent:\t", r.UserAgent())

	ua := user_agent.New(r.UserAgent())

	fmt.Printf("is mobile:\t%v\n", ua.Mobile()) // => false
	fmt.Printf("is robot:\t%v\n", ua.Bot())     // => false
	fmt.Printf("mozilla:\t%v\n", ua.Mozilla())  // => "5.0"

	fmt.Printf("platform:\t%v\n", ua.Platform()) // => "X11"
	fmt.Printf("system:\t%v\n", ua.OS())         // => "Linux x86_64"

	name, version := ua.Engine()
	fmt.Printf("browser kernel:\t%v\n", name)            // => "AppleWebKit"
	fmt.Printf("browser kernel version:\t%v\n", version) // => "537.11"

	name, version = ua.Browser()
	fmt.Printf("browser:\t%v\n", name)            // => "Chrome"
	fmt.Printf("browser version:\t%v\n", version) // => "23.0.1271.97"

	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
}

func main() {
	http.HandleFunc("/", showInfo)
	http.HandleFunc("/visit", tongji)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
