package data

import (
	"fmt"
	"net/http"
)

// 资源结构
type Resource map[string]string

// 资源接口
type ResourceFunc func(Resource, *http.Request)

var dataSource = make(map[string]ResourceFunc)

func GetResource(r *http.Request) string {
	source := make(Resource)

	for _, v := range dataSource {
		v(source, r)
	}

	result := ""

	for k, v := range source {
		result += k + ":" + v + "\r\n"
	}

	result += "\r\n"

	return result
}

// 注册函数，将资源接口注册进列表
func Register(name string, r ResourceFunc) {
	if _, exists := dataSource[name]; !exists {
		dataSource[name] = r
	} else {
		fmt.Println("Resource exists!")
	}
}
