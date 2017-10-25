package config

import (
	"time"
)

type Config struct {
	KafkaServers   []string
	KafkaTopic     string
	TimeLocation   *time.Location
	MessageBuffNum int
}

var CommonConfig Config
