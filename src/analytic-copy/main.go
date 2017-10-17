package main

import (
	"analytic-copy/data"
	"analytic-copy/db"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var countdown uint32 = 0
var beforeCountdown uint32 = 0

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint32(&countdown, 1)

	var result = data.GetResource(r)

	db.ProducerOne(result)

	fmt.Fprintf(w, "ok")
}

func main() {
	// 设置路由规则
	http.HandleFunc("/page-visit", pageVisit)

	fmt.Println("start server")

	go func() {
		for {
			<-time.After(time.Second)
			fmt.Printf("%s:\ttotal request: %d\tqps: %d\n\n", time.Now().Format("2006-01-02 15:04:05"), countdown, countdown-beforeCountdown)
			beforeCountdown = countdown
		}
	}()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
