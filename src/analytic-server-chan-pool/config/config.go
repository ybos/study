package config

import (
	"time"

	"github.com/Shopify/sarama"
)

type Config struct {
	KafkaServers   []string
	KafkaTopic     string
	TimeLocation   *time.Location
	MessageBuffNum int
	MaxWorker      int
}

var CommonConfig Config

func NewKafkaConfig() *sarama.Config {
	conf := sarama.NewConfig()

	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	//	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0

	// producer.Successes() 是一个成功消息的通知管道，是个有缓存通道
	// producer.Errors() 是一个失败消息的通知管道，是个有缓存通道
	// 只有设置把 Return.Successes 设置成 true 才可以访问通道，如果不及时访问清理消息，可能造成阻塞
	conf.Producer.Return.Successes = true

	return conf
}
