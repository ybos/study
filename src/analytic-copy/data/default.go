package data

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// 资源结构
type Resource = bytes.Buffer

// 资源接口
type ResourceFunc func(*Resource, *http.Request)

var dataSource = make(map[string]ResourceFunc)

func GetResource(r *http.Request) string {
	start_time := time.Now()
	//	source := make(Resource)
	var source Resource

	for _, v := range dataSource {
		v(&source, r)
	}

	source.WriteString("\r\n")

	fmt.Println("use time: ", time.Since(start_time))

	return source.String()
}

// 注册函数，将资源接口注册进列表
func Register(name string, r ResourceFunc) {
	if _, exists := dataSource[name]; !exists {
		dataSource[name] = r
	} else {
		fmt.Println("Resource exists!")
	}
}
