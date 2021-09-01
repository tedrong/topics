package crontab

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/topics/crawler"
	"github.com/topics/models"
)

type Highlight struct {
	BasicCron
}

var HighlightModel = new(models.Highlight)

func (m *Highlight) Period() string {
	// return "@hourly"
	return "* * * * *"
}

func (m *Highlight) Do() {
	crawler := crawler.Get()
	crawler.Mutex.Lock()
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/exchange/FMTQIK.html"
	crawler.GOTO()
	time.Sleep(2 * time.Second)
	// Find the latest record in database, return 1970-01-01 if empty
	date := HighlightModel.LatestDate()
	log.Printf("The latest date of Highlight is %s", date)
	// Startup crawler with date(first day of current month)
	Highlight, err := crawler.Highlight(time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Print(errors.Wrap(err, "Get highlight fail"))
	}
	HighlightModel.Store(Highlight)
	crawler.Mutex.Unlock()
}
