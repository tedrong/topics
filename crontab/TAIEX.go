package crontab

import (
	"fmt"
	"time"

	"github.com/topics/crawler"
	"github.com/topics/logging"
	"github.com/topics/models"
)

type TAIEX struct {
	BasicCron
}

var TAIEXModel = new(models.TAIEXModel)

func (t *TAIEX) Period() string {
	// return "@hourly"
	return "* * * * *"
}

func (t *TAIEX) Do() {
	crawler := crawler.Get()
	crawler.Mutex.Lock()
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/indices/MI_5MINS_HIST.html"
	crawler.GOTO()
	time.Sleep(2 * time.Second)
	// Find the latest record in database, return 1970-01-01 if empty
	date := TAIEXModel.LatestDate()
	t.LogJob(logging.Get().Info(), CJ_TAIEX).Msg(fmt.Sprintf("The latest date of TAIEX is %s", date))
	// Startup crawler with date(first day of current month)
	TAIEX, err := crawler.TAIEX(time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.LogJob(logging.Get().Warn(), CJ_TAIEX).Err(err)
	}
	TAIEXModel.Store(TAIEX)
	crawler.Mutex.Unlock()
}
