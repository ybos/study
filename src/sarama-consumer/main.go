package main

import (
	"log"
	"runtime"
	"strings"
	//	"time"

	"github.com/Shopify/sarama"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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

	defer consumer.Close()
	defer client.Close()

	for {
		select {
		case err := <-client.Errors():
			log.Println("haha: ", err.Error())
		case msg := <-client.Messages():
			log.Printf("Partition:%d, Offset:%d\n", msg.Partition, msg.Offset)
		}
		//		<-time.After(1 * time.Millisecond)
	}

	log.Println("+++++++++++++++++++")
}
