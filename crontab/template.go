package crontab

import (
	"encoding/json"
	"log"
)

var (
	CJ_Trends = "GoogleTrends"
)

type BasicCron struct{}

func (b *BasicCron) LogJob(jobName string, logContent interface{}) {
	msg := map[string]interface{}{
		"type":    "cron job",
		"server":  "topics",
		"CronJob": jobName,
		"Log":     logContent,
	}

	byteString, err := json.Marshal(msg)
	if err == nil {
		log.Println(string(byteString))
	}

}
