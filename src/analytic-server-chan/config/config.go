package config

import (
	"time"
)

type Config struct {
	KafkaServers []string
	KafkaTopic   string
	TimeLocation *time.Location
}

var CommonConfig Config
