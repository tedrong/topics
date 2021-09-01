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
		log.Print(errors.Wrap(err, "Get daily trading fail"))
	}
	DailyTradingModel.Store(DailyTrading)

	// Get stocks supplement information from the other page (Stock P/E ratio, dividend yield and P/B ratio)
	crawler.URL = "https://www.twse.com.tw/zh/page/trading/exchange/BWIBBU_d.html"
	crawler.GOTO()
	// Find the latest record in database, return 1970-01-01 if empty
	date := DailyTradingModel.LatestRatioDate()
	log.Printf("The latest date of daily trading supplement is %s", date)

	time.Sleep(2 * time.Second)
	DailyTrading, err = crawler.DailyTradingRatio(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Print(errors.Wrap(err, "Get daily trading supplement fail"))
	}
	DailyTradingModel.Store(DailyTrading)

	crawler.Mutex.Unlock()
}
