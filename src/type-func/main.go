package main

import (
	. "fmt"
)

// 这是一个定义有 Add 方法的接口
type Analytics interface {
	Add(string) string
}

// ###########################
type Website func(domain string) string

func (w Website) Add(path string) string {
	return w(path) + "/?username=administrator"
}

// ###########################
func process(a Analytics) {
	Println("process: ", a.Add("ybo.me"))
}

// ###########################
func anotherFunc(path string) string {
	return path
}

func main() {
	var website1 Website = func(domain string) string {
		return domain + "-" + domain
	}

	Println(website1("ybo.me"))
	Println(website1.Add("ybo.me"))

	process(website1)
	process(Website(anotherFunc))
}
