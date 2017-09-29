package data

import (
	//	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mssola/user_agent"
)

func init() {
	Register("browser", Get)
}

// 获取通过网络请求来获取的基本信息
func Get(data Resource, r *http.Request) {
	data["server_time"] = time.Now().Local().String()
	data["server_client_ip"] = r.RemoteAddr
	data["param_sessionid"] = r.FormValue("sessionid")
	data["param_userid"] = r.FormValue("userid")

	data["param_url"] = r.FormValue("url")

	urlData, err := url.Parse(data["param_url"])
	if err == nil {
		data["param_url_hostname"] = urlData.Hostname()
	} else {
		data["param_url_hostname"] = ""
	}

	data["param_referer"] = r.FormValue("referer")

	urlData, err = url.Parse(data["param_referer"])
	if err == nil {
		data["param_referer_hostname"] = urlData.Hostname()
	} else {
		data["param_referer_hostname"] = ""
	}

	ua := user_agent.New(r.UserAgent())

	data["server_ua_platform"] = ua.Platform()
	data["server_ua_system"] = ua.OS()

	name, version := ua.Browser()
	data["server_ua_browser_name"] = name
	data["server_ua_browser_version"] = version

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

	//	ua = user_agent.New(r.UserAgent())

	//	fmt.Printf("operating platform:\t%v\n", ua.Platform())
	//	fmt.Printf("operating system:\t%v\n", ua.OS())

	//	name, version = ua.Browser()
	//	fmt.Printf("browser name:\t%v\n", name)
	//	fmt.Printf("browser version:\t%v\n", version)

	//	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
}
