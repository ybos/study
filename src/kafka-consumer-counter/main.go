package main

import (
	"time"
	//	"flag"
	"fmt"
	//	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/Shopify/sarama"
	//	"github.com/bsm/sarama-cluster"
)

func main() {
	//	var err error
	//	var _startTime string
	//	var _endTime string

	//	flag.StringVar(&_startTime, "start_time", "", "put your start_time here")
	//	flag.StringVar(&_endTime, "end_time", "", "put your end_time here")

	//	flag.Parse()

	//	if flag.NFlag() != 2 {
	//		fmt.Println("请使用参数 -start_time 和 -end_time")
	//		return
	//	}

	//	startTime, err := time.Parse("2006-01-02 15:04:05", _startTime)
	//	if err != nil {
	//		fmt.Println("转换开始时间时出错, 错误日志为: ", err.Error())
	//		return
	//	}

	//	endTime, err := time.Parse("2006-01-02 15:04:05", _endTime)
	//	if err != nil {
	//		fmt.Println("转换结束时间时出错, 错误日志为: ", err.Error())
	//		return
	//	}

	//	if startTime.Unix() > endTime.Unix() {
	//		fmt.Println("开始时间不能大于结束时间")
	//		return
	//	}

	var consumerTopic string = "test"
	//	var consumerGroupID string = "test-group-1"
	//	config := cluster.NewConfig()
	//	config.Group.Return.Notifications = true
	//	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	//	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 初始从最新的offset开始

	//	c, err := cluster.NewConsumer(strings.Split("10.10.11.13:9092", ","), consumerGroupID, strings.Split(consumerTopic, ","), config)

	config := sarama.NewConfig()
	// 开发环境中我们使用 sarama.OffsetOldest，Kafka 将从创建以来第一条消息开始发送。
	// 在生产环境中切换为 sarama.OffsetNewest, 只会将最新生成的消息发送给我们。
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	c, err := sarama.NewConsumer(strings.Split("10.10.11.13:9092", ","), config)

	consumer, err := c.ConsumePartition(consumerTopic, 0, sarama.OffsetOldest)

	if err != nil {
		fmt.Println("failed to open consumer: ", err.Error())
		return
	}

	defer c.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	var countdown uint32 = 0

	// 处理错误的
	go func(c sarama.PartitionConsumer, countdown *uint32) {
		for {
			select {
			case err := <-c.Errors():
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++", err.Error())
			case <-c.Messages():
				//				fmt.Printf("%s/%d/%d\t%s\n\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
				atomic.AddUint32(countdown, 1)
			}
		}

		fmt.Println("================================")
	}(consumer, &countdown)

	go func(countdown *uint32) {
		for {
			<-time.After(time.Second)
			fmt.Println(*countdown)
		}
	}(&countdown)

	//	fmt.Println("准备开始：")

	//	for msg := range consumer.Messages() {
	//		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
	//		fmt.Println("------------------------------")
	//		//		c.MarkOffset(msg, "") // MarkOffset 并不是试试写入 kafka，有可能在程序 crash 时丢失未提交的 offset
	//	}

	wg.Wait()
}
