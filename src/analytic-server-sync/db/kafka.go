package db

import (
	"fmt"
	"os"

	"analytic-server-sync/config"

	"github.com/Shopify/sarama"
)

var KafkaProducer sarama.SyncProducer

var KafkaChannel chan string

// 初始化 kafka producer
// 销毁 kafka producer
// 设置共享的 kafka producer
func CreateKafkaProducer() {
	var err error

	KafkaProducer, err = sarama.NewSyncProducer(config.CommonConfig.KafkaServers, config.NewKafkaConfig())

	if err != nil {
		fmt.Println("Failed to create producer: ", err.Error())
		os.Exit(1)
	}

	KafkaChannel = make(chan string, config.CommonConfig.MessageBuffNum)

	go AsyncProduceMessage()
}

// 将消息写入一个缓存
// 这一步有待考察, 通道是个有锁类型, 锁的开销在程序中占比会非常大
// 可以考虑将这一步换做自定义的RWMutex配合批量写入(如果kafka支持的话)
func ProducerOne(s string) {
	KafkaChannel <- s
}

// 异步提交消息
func AsyncProduceMessage() {
	for {
		s := <-KafkaChannel

		// 消息的结构
		msg := &sarama.ProducerMessage{
			Topic: config.CommonConfig.KafkaTopic,
			Value: sarama.StringEncoder(s),
		}

		// 写入消息
		_, _, err := KafkaProducer.SendMessage(msg)

		if err != nil {
			fmt.Printf("Kafka error: %s\n", err.Error())
		}
	}
}
