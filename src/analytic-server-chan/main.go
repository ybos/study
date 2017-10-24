package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"analytic-chan-eg/config"
	"analytic-chan-eg/data"
	"analytic-chan-eg/db"
)

func init() {
	// 初始化设置
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化时间相关
	config.CommonConfig.TimeLocation, _ = time.LoadLocation("Asia/Shanghai")

	// 初始化参数
	var _kafkaServers, _kafkaTopic string

	flag.StringVar(&_kafkaServers, "kafka-servers", "10.10.11.13:9092", "请输入Kafka服务器地址(server1:9092,server2:9092,server3:9092)")
	flag.StringVar(&_kafkaTopic, "kafka-topic", "test", "请输入Kafka的话题名(example-topic)")

	flag.Parse()

	if _kafkaServers == "" {
		fmt.Println("请输入Kafka服务器地址(server1:9092,server2:9092,server3:9092)")
		os.Exit(0)
	}

	config.CommonConfig.KafkaServers = strings.Split(_kafkaServers, ",")

	if _kafkaTopic == "" {
		fmt.Println("请输入Kafka的话题名(example-topic)")
		os.Exit(0)
	}

	config.CommonConfig.KafkaTopic = _kafkaTopic

	// 初始化数据库
	db.CreateKafkaProducer()
}

// 访问统计的实现函数
func pageVisit(w http.ResponseWriter, r *http.Request) {
	var result = data.GetResource(r)

	db.ProducerOne(result)

	w.WriteHeader(http.StatusOK)
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
