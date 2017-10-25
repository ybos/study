package db

import (
	"fmt"
	"os"
	"time"

	"analytic-server-chan/config"

	"github.com/Shopify/sarama"
)

var KafkaProducer sarama.AsyncProducer

var KafkaChannel chan string = make(chan string, config.CommonConfig.MessageBuffNum)

// 初始化 kafka producer
// 销毁 kafka producer
// 设置共享的 kafka producer
func CreateKafkaProducer() {
	// 创建新的配置
	saramaConfig := sarama.NewConfig()
	// producer.Successes() 是一个成功消息的通知管道，是个有缓存通道
	// producer.Errors() 是一个失败消息的通知管道，是个有缓存通道
	// 只有设置把 Return.Successes 设置成 true 才可以访问通道，如果不及时访问清理消息，可能造成阻塞
	saramaConfig.Producer.Return.Successes = true
	// 设置超时时间
	saramaConfig.Producer.Timeout = 5 * time.Second

	// 创建一个写对象
	// SyncProducer 是对 ASyncProducer 的一个封装
	// SyncProducer 每发送一条消息就要等待返回，所以不可以异步
	var err error
	KafkaProducer, err = sarama.NewAsyncProducer(config.CommonConfig.KafkaServers, saramaConfig)

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
			fmt.Println("Failed to insert msg: ", err.Error())
		case <-successes:
		}
	}
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
		KafkaProducer.Input() <- msg
	}
}
