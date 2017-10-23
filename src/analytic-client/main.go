package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// 初始化设置
	runtime.GOMAXPROCS(runtime.NumCPU())
	local, _ := time.LoadLocation("Asia/Shanghai")

	// 初始化参数
	var _startTime, _endTime string
	var startTime, endTime int64

	flag.StringVar(&_startTime, "start-time", "", "输入您需要查询的起始时间(2006-01-02 15:04:05)")
	flag.StringVar(&_endTime, "end-time", "", "输入您需要查询的结束时间(2006-01-02 15:04:05)")

	flag.Parse()

	if _startTime != "" {
		_startTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", _startTime, local)
		if err != nil {
			fmt.Println("转换开始时间时出错, 错误提示为: ", err.Error())
			return
		}

		startTime = _startTimeObj.Unix()
	} else {
		startTime = 0
	}

	if _endTime != "" {
		_endTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", _endTime, local)
		if err != nil {
			fmt.Println("转换结束时间时出错, 错误提示为: ", err.Error())
			return
		}

		endTime = _endTimeObj.Unix()
	} else {
		endTime = time.Now().Local().Unix()
	}

	if startTime > endTime {
		fmt.Println("查询的起始时间不能大于结束时间")
		return
	}

	fmt.Println("您需要查询的时间范围为：", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"), " 至 ", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("如需更改时间请使用 --start-time 和 --end-time 参数")

	// Kafka 基本信息获取
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	connect, err := sarama.NewConsumer(strings.Split("10.10.11.13:9092", ","), config)

	if err != nil {
		fmt.Println("创建一个连接失败: ", err.Error())
		return
	}

	oldestConsumer, err := connect.ConsumePartition("test", 0, sarama.OffsetOldest)

	if err != nil {
		fmt.Println("创建最老的数据消费者失败: ", err.Error())
		return
	}

	defer oldestConsumer.Close()

	client, err := sarama.NewClient(strings.Split("10.10.11.13:9092", ","), nil)

	if err != nil {
		fmt.Println("创建实例失败: ", err.Error())
	}

	defer client.Close()

	maxOffset, err := client.GetOffset("test", 0, sarama.OffsetNewest)

	if err != nil {
		fmt.Println("获取数据偏移量失败: ", err.Error())
	}

	var countPageVisit uint32
	var countRecord int64

Out:
	for {
		select {
		case err := <-oldestConsumer.Errors():
			fmt.Println("数据获取失败: ", err.Error())
			break Out
		case msg := <-oldestConsumer.Messages():
			countRecord++

			data := parseRecord(string(msg.Value))

			serverTime, err := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", data["server_time"], local)

			if err != nil {
				continue
			}

			if serverTime.Unix() >= startTime && serverTime.Unix() < endTime {
				countPageVisit++
			} else if serverTime.Unix() >= endTime {
				break Out
			}
		default:
			if countRecord >= maxOffset {
				break Out
			}
		}
	}

	fmt.Println("总数据量: ", maxOffset)
	fmt.Println("查询数量: ", countRecord)
	fmt.Println("查询结果: ", countPageVisit)
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
