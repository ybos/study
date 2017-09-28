package main

import (
	"log"
	"os"
	"strings"
	"sync"

	"kafka-base/sarama"
)

var (
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)
)

func ProduceMessage() {
	sarama.Logger = logger

	config := sarama.NewConfig()
	// SyncProducer 是对 AsyncProducer 作的封装
	// SyncProducer每发送出去一条消息，就等待返回结果，然后再发下一条。因此SyncProducer不支持批量发送。
	// AsyncProducer的Input()和Successes()的阻塞问题
	// AsyncProducer的两个方法：Input()返回用来写入消息的channel，Successes()返回用来收集发送成功的消息的channel(Errors()用来收集发送失败的消息)。
	// 应用程序可以用一个goroutine不断地向Input()写入消息，用另一个goroutine从Successes()和Errors()里读取发送结果，以此实现异步发送。
	// 只有sarama.Config.Producer.Return.Successes设置为true，才可以从producer.Successes()里读取。
	// 而且，如果该参数为true，必须读取producer.Successes()
	// 否则producer.successes channel就满，进而导致producer.input channel也满，然后写producer.Input()的时候就阻塞了。
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("你好,世界!")

	// 同步的
	producer, err := sarama.NewSyncProducer(strings.Split("10.10.11.13:9092", ","), config)

	if err != nil {
		logger.Println("Failed to produce message: ", err)

		os.Exit(1)
	}

	defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		logger.Println("Failed to produce msg: ", err)

		os.Exit(1)
	}

	logger.Printf("partition=%d, offset=%d\n", partition, offset)
}

func main() {
	ProduceMessage()
}
