package server

import (
	"analytic-server-chan-pool/config"
	"analytic-server-chan-pool/data"
	"analytic-server-chan-pool/db"

	"fmt"
	"net/http"
)

var counter = make(chan struct{}, config.CommonConfig.MaxWorker*100)

func releaseCounter() {
	<-counter
}

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	counter <- struct{}{}

	defer releaseCounter()

	var result = data.GetResource(r)

	db.ProducerOne(result)

	w.WriteHeader(http.StatusOK)
}

func CreateServer() {
	// 设置路由规则
	http.HandleFunc("/page-visit", pageVisit)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
	}
}
