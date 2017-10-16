package db

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

const _kafka_topic string = "wuliang"

var KafkaProducer sarama.AsyncProducer

// 初始化 kafka producer
// 销毁 kafka producer
// 设置共享的 kafka producer
func init() {
	// 创建新的配置
	config := sarama.NewConfig()
	// producer.Successes() 是一个成功消息的通知管道，是个有缓存通道
	// producer.Errors() 是一个失败消息的通知管道，是个有缓存通道
	// 只有设置把 Return.Successes 设置成 true 才可以访问通道，如果不及时访问清理消息，可能造成阻塞
	config.Producer.Return.Successes = true
	// 设置超时时间
	config.Producer.Timeout = 5 * time.Second

	// 创建一个写对象
	// SyncProducer 是对 ASyncProducer 的一个封装
	// SyncProducer 每发送一条消息就要等待返回，所以不可以异步
	var err error
	KafkaProducer, err = sarama.NewAsyncProducer(strings.Split("10.10.11.13:9092", ","), config)

	if err != nil {
		fmt.Println("Failed to create producer: ", err.Error())
		os.Exit(1)
	}

	go consumeMessage(KafkaProducer)
}

// 及时的将通道清空, 防止阻塞
func consumeMessage(p sarama.AsyncProducer) {
	// 消息写入成功后关闭对象
	defer p.Close()

	errors := p.Errors()
	successes := p.Successes()

	for {
		select {
		case err := <-errors:
			fmt.Println("Failed to send a message: ", err.Error())
		case <-successes:
		}
	}
}

func ProducerOne(p sarama.AsyncProducer, s string) {
	// 消息的结构
	msg := &sarama.ProducerMessage{
		Topic: _kafka_topic,
		Value: sarama.ByteEncoder(s),
	}

	// 写入消息
	p.Input() <- msg
}
