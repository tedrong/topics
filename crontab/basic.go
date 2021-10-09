package crontab

import (
	"github.com/rs/zerolog"
)

var (
	CJ_Trends       = "GoogleTrends"
	CJ_DailyTrading = "StockDailyTrading"
	CJ_Highlight    = "Highlight"
	CJ_TAIEX        = "TAIEX"
)

type BasicCron struct{}

func (b *BasicCron) LogJob(zlog *zerolog.Event, jobName string) *zerolog.Event {
	zlog.
		Str("type", "CronTab").
		Str("server", "Topics").
		Str("cronjob", jobName)
	return zlog
}
