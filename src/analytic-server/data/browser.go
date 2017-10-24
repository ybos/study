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
func Get(data *Resource, r *http.Request) {
	data.WriteString("server_time:" + time.Now().Local().String() + "\r\n")
	data.WriteString("server_client_ip:" + r.RemoteAddr + "\r\n")
	ua := user_agent.New(r.UserAgent())

	data.WriteString("server_ua_platform:" + ua.Platform() + "\r\n")
	data.WriteString("server_ua_system:" + ua.OS() + "\r\n")

	if ua.Mobile() {
		data.WriteString("server_ua_is_mobile:" + "true" + "\r\n")
	} else {
		data.WriteString("server_ua_is_mobile:" + "false" + "\r\n")
	}

	name, version := ua.Browser()
	data.WriteString("server_ua_browser_name:" + name + "\r\n")
	data.WriteString("server_ua_browser_version:" + version + "\r\n")

	data.WriteString("param_sessionid:" + r.FormValue("sessionid") + "\r\n")
	data.WriteString("param_userid:" + r.FormValue("userid") + "\r\n")

	data.WriteString("param_url:" + r.FormValue("url") + "\r\n")

	urlData, err := url.Parse(r.FormValue("url"))
	if err == nil {
		data.WriteString("param_url_hostname:" + urlData.Hostname() + "\r\n")
	} else {
		data.WriteString("param_url_hostname:" + "" + "\r\n")
	}

	data.WriteString("param_referer:" + r.FormValue("referer") + "\r\n")

	urlData, err = url.Parse(r.FormValue("referer"))
	if err == nil {
		data.WriteString("param_referer_hostname:" + urlData.Hostname() + "\r\n")
	} else {
		data.WriteString("param_referer_hostname:" + "" + "\r\n")
	}
}
