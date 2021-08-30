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

var DailyTradingModel = new(models.DailyTrading)
var StockInfoModel = new(models.StockInfoModel)

func (m *DailyTrading) Period() string {
	// return "@hourly"
	return "* * * * *"
}

func (m *DailyTrading) Do() {
	crawler := crawler.Get()
	crawler.Mutex.Lock()
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/exchange/STOCK_DAY.html"
	crawler.GOTO()
	time.Sleep(2 * time.Second)
	// Get all stocks from datbase, we'll use the symbol code to fetch individual information for each stock.
	stocks := StockInfoModel.All()
	// Startup crawler with date(first day of current month)
	DailyTrading, err := crawler.DailyTrading(stocks)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Get daily trading fail"))
	}
	DailyTradingModel.Store(DailyTrading)
	crawler.Mutex.Unlock()
}
