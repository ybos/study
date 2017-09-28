package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

const topics string = "test"

func syncProducer() {
	// 创建新的配置
	config := sarama.NewConfig()
	// producer.Successes() 是一个成功消息的通知管道，是个有缓存通道
	// producer.Errors() 是一个失败消息的通知管道，是个有缓存通道
	// 只有设置把 Return.Successes 设置成 true 才可以访问通道，如果不及时访问清理消息，可能造成阻塞
	config.Producer.Return.Successes = true
	// 设置超时时间
	config.Producer.Timeout = 5 * time.Second

	// 创建一个写对象
	producer, err := sarama.NewAsyncProducer(strings.Split("10.10.11.13:9092", ","), config)

	defer producer.Close()

	if err != nil {
		log.Println("Failed to create producer: ", err.Error())
		os.Exit(1)
	}

	// 处理成功/错误消息,防止阻塞
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		successes := p.Successes()

		for {
			select {
			case err := <-errors:
				log.Println("Failed to send a message: ", err.Error())
			case <-successes:
			}
		}
	}(producer)

	messageContent := "async: " + strconv.Itoa(time.Now().UnixNano())

	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(messageContent),
	}

	producer.Input() <- msg
}
