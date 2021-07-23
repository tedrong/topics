package crontab

import (
	"log"
	"time"
	"topics/config"
	"topics/crawler"
	"topics/database"
	"topics/sysexec"

	"github.com/pkg/errors"
)

type MarketIndex struct {
	BasicCron
}

func (m *MarketIndex) Period() string {
	// return "@hourly"
	return "0 * * * *"
}

func (m *MarketIndex) Do() {
	// Get system configs for crawler
	cfg := config.Get()

	// Check if there is a instance running, kill it
	if pid := sysexec.FindWebDriverPID(); pid != nil {
		sysexec.KillWebDriver(pid)
	}

	// Initialize
	crawlerEntry := crawler.CrawlerEntry{
		URL:             "https://www.twse.com.tw/zh/page/trading/indices/MI_5MINS_HIST.html",
		SeleniumPath:    cfg.Crawler.SeleniumPath,
		GeckoDriverPath: cfg.Crawler.GeckoDriverPath,
		Port:            cfg.Crawler.Port,
	}
	// Driver instance startup
	webDriver, err := crawlerEntry.StartWebInstance()
	if err != nil {
		log.Fatal(errors.Wrap(err, "WebDriver instance startup fail"))
	}
	// New crawler with URL
	crawlerEntry.Crawler, err = crawlerEntry.Init()
	if err != nil {
		log.Fatal(errors.Wrap(err, "URL connection fail"))
	}
	// Create new data transfer object
	dto := database.MarketIndexDTO{
		BasicDTO:    database.BasicDTO{},
		MarketIndex: database.MarketIndex{},
	}
	// Find the latest record in database, return 1970-01-01 if empty
	date := dto.LatestDate()
	log.Printf("The latest date of market index is %s", date)
	// Startup crawler with date(first day of current month)
	markets, err := crawlerEntry.MarketIndex(time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Fatal(errors.Wrap(err, "Get market indexes fail"))
	}
	// Create new row in database
	for _, element := range *markets {
		dto.MarketIndex = *element
		dto.Insert()
	}
	// Stop crawler and web driver
	defer (*crawlerEntry.Crawler).Quit()
	defer webDriver.Stop()
}
