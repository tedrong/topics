package crontab

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/topics/crawler"
	"github.com/topics/models"
)

type TAIEX struct {
	BasicCron
}

var TAIEXModel = new(models.TAIEXModel)

func (m *TAIEX) Period() string {
	// return "@hourly"
	return "* * * * *"
}

func (m *TAIEX) Do() {
	crawler := crawler.Get()
	crawler.Mutex.Lock()
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/indices/MI_5MINS_HIST.html"
	crawler.GOTO()
	time.Sleep(2 * time.Second)
	// Find the latest record in database, return 1970-01-01 if empty
	date := TAIEXModel.LatestDate()
	log.Printf("The latest date of TAIEX is %s", date)
	// Startup crawler with date(first day of current month)
	TAIEX, err := crawler.TAIEX(time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Fatal(errors.Wrap(err, "Get TAIEX fail"))
	}
	TAIEXModel.Store(TAIEX)
	crawler.Mutex.Unlock()
}
