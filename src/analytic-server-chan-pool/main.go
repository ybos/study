package main

import (
	"flag"
	"fmt"
	//	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"

	"analytic-server-chan-pool/config"
	"analytic-server-chan-pool/db"
	"analytic-server-chan-pool/server"
	"analytic-server-chan-pool/worker"
)

func init() {
	// 初始化设置
	var cpuNum = runtime.NumCPU()

	config.CommonConfig.MaxWorker = (cpuNum / 2) + 1

	runtime.GOMAXPROCS(cpuNum)

	// 初始化时间相关
	config.CommonConfig.TimeLocation, _ = time.LoadLocation("Asia/Shanghai")

	// 初始化参数
	var _kafkaServers, _kafkaTopic string
	var _messageBuffNum int

	flag.StringVar(&_kafkaServers, "kafka-servers", "", "请输入Kafka服务器地址(server1:9092,server2:9092,server3:9092)")
	flag.StringVar(&_kafkaTopic, "kafka-topic", "", "请输入Kafka的话题名(example-topic)")
	flag.IntVar(&_messageBuffNum, "message-buff-num", 10000, "请输入消息缓存数量(默认10000)")

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

	config.CommonConfig.MessageBuffNum = _messageBuffNum
}

func main() {
	fmt.Println("start server")

	// 初始化数据库
	db.CreateKafkaProducer()

	// 初始化处理器
	worker.CreateWorker()

	// 初始化服务器
	server.CreateServer()
}
