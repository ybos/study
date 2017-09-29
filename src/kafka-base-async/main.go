package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

const topics string = "test"

var wg sync.WaitGroup

func syncProducer() {
	// 创建新的配置
	config := sarama.NewConfig()
	// producer.Successes() 是一个成功消息的通知管道，是个有缓存通道
	// producer.Errors() 是一个失败消息的通知管道，是个有缓存通道
	// 只有设置把 Return.Successes 设置成 true 才可以访问通道，如果不及时访问清理消息，可能造成阻塞
	config.Producer.Return.Successes = true
	// 设置超时时间
	config.Producer.Timeout = 5 * time.Second
	// 设置最多允许多少条消息缓存
	//	config.ChannelBufferSize = 1

	// 创建一个写对象
	// SyncProducer 是对 ASyncProducer 的一个封装
	// SyncProducer 每发送一条消息就要等待返回，所以不可以异步
	producer, err := sarama.NewAsyncProducer(strings.Split("10.10.11.13:9092", ","), config)

	// 消息写入成功后关闭对象
	defer producer.Close()

	if err != nil {
		log.Println("Failed to create producer: ", err.Error())
		os.Exit(1)
	}

	// 这次加的是单次 goroutine  函数
	wg.Add(1)

	// 处理成功/错误消息,防止阻塞
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		successes := p.Successes()

	Wait:
		for {
			select {
			case err := <-errors:
				log.Println("Failed to send a message: ", err.Error())
			case <-successes:
				log.Println("Successed to send a message on ", time.Now())
			case <-time.After(10 * time.Second):
				break Wait
			}
		}

		wg.Done()
	}(producer)

	// 开始尝试批量发送消息
	var concurrentNum int = 100
	wg.Add(concurrentNum)

	// 5秒后的时间点
	goalTime := time.Now().Add(5 * time.Second)
	// 时间区间
	duration := time.Since(goalTime)

	for i := 0; i < concurrentNum; i++ {
		go func(d time.Duration, p sarama.AsyncProducer) {
			<-time.After(d)

			// 消息的内容
			messageContent := "async: " + time.Now().String()

			// 消息的结构
			msg := &sarama.ProducerMessage{
				Topic: topics,
				Value: sarama.ByteEncoder(messageContent),
			}

			// 写入消息
			p.Input() <- msg

			// 解除一个任务
			wg.Done()
		}(duration, producer)
	}

	wg.Wait()
}

func main() {
	syncProducer()
}
