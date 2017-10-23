package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	local, _ := time.LoadLocation("Asia/Shanghai")

	var _startTime, _endTime string

	flag.StringVar(&_startTime, "start-time", "", "put your start time here(2006-01-02 15:04:05)")
	flag.StringVar(&_endTime, "end-time", "", "put your end time here(2006-01-02 15:04:05)")

	flag.Parse()

	if flag.NFlag() != 2 {
		fmt.Println("请使用参数 -start-time 和 -end-time")
		return
	}

	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", _startTime, local)
	if err != nil {
		fmt.Println("转换开始时间时出错, 错误日志为: ", err.Error())
		return
	}

	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", _endTime, local)
	if err != nil {
		fmt.Println("转换结束时间时出错, 错误日志为: ", err.Error())
		return
	}

	startTimestamp := startTime.Unix()
	endTimestamp := endTime.Unix()

	if startTimestamp > endTimestamp {
		fmt.Println("开始时间不能大于结束时间")
		return
	}

	fmt.Printf("正在统计[start-time, end-time)中的数据...\n")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// 开发环境中我们使用 sarama.OffsetOldest，Kafka 将从创建以来第一条消息开始发送。
	// 在生产环境中切换为 sarama.OffsetNewest, 只会将最新生成的消息发送给我们。
	//	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(strings.Split("10.10.11.13:9092", ","), config)

	if err != nil {
		log.Println("failed to create consumer: ", err.Error())
		return
	}

	client, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)

	if err != nil {
		log.Println("failed to open consumer: ", err.Error())
		return
	}

	fmt.Println(time.Now().In(local))

	defer consumer.Close()
	defer client.Close()

	var countPageVisit uint32

Out:
	for {
		select {
		case err := <-client.Errors():
			log.Println("haha: ", err.Error())
			break Out
		case msg := <-client.Messages():
			data := parseRecord(string(msg.Value))

			serverTime, err := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", data["server_time"], local)

			if err != nil {
				continue
			}

			if serverTime.Unix() >= startTimestamp && serverTime.Unix() < endTimestamp {
				countPageVisit++
			} else if serverTime.Unix() >= endTimestamp {
				break Out
			}
		}
	}

	log.Println("page visit: ", countPageVisit, "\n+++++++++++++++++++")
}

func parseRecord(record string) map[string]string {
	data := strings.Split(strings.Trim(record, "\r\n"), "\r\n")

	dataMap := make(map[string]string)
	for _, v := range data {
		d := strings.SplitN(v, ":", 2)
		dataMap[d[0]] = d[1]
	}

	return dataMap
}
