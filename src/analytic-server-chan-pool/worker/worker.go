package worker

import (
	"analytic-server-chan-pool/config"
	"analytic-server-chan-pool/db"
)

func CreateWorker() {
	for i := 0; i < config.CommonConfig.MaxWorker; i++ {
		go db.AsyncProduceMessage()
	}
}
