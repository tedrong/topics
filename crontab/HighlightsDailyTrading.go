package crontab

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/topics/crawler"
	"github.com/topics/models"
)

type DailyTrading struct {
	BasicCron
}

var DailyTradingModel = new(models.DailyTradingModel)

func (m *DailyTrading) Period() string {
	// return "@hourly"
	return "* * * * *"
}

func (m *DailyTrading) Do() {
	crawler := crawler.Get()
	crawler.Mutex.Lock()
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/exchange/FMTQIK.html"
	crawler.GOTO()
	time.Sleep(2 * time.Second)
	// Find the latest record in database, return 1970-01-01 if empty
	date := DailyTradingModel.LatestDate()
	log.Printf("The latest date of DailyTrading is %s", date)
	// Startup crawler with date(first day of current month)
	HighlightsDailyTrading, err := crawler.DailyTrading(time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Fatal(errors.Wrap(err, "Get market indexes fail"))
	}
	DailyTradingModel.Store(HighlightsDailyTrading)
	crawler.Mutex.Unlock()
}
