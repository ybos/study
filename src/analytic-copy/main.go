package main

import (
	"analytic-copy/data"
	"analytic-copy/db"
	"fmt"
	"net/http"
)

//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  500.1µs
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  499.5µs
//use time:  0s
//use time:  0s
//use time:  0s
//use time:  0s

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	var result = data.GetResource(r)

	db.ProducerOne(db.KafkaProducer, result)

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
