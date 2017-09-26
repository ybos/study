package main

import (
	"fmt"
	"net/http"
	//	"regexp"
	"time"

	"http-request-info/user_agent"
)

//type MachineInfo struct {
//	Browser  string
//	System   string
//	Kernel   string
//	Language string
//}

//type ParseMachineInfo(userAgent string) MachineInfo {
//	m, _ := regexp.MatchString("", userAgent)

//	fmt.Println(m)
//}

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

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
