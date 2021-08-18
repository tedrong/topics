package crontab

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/groovili/gogtrends"
	"github.com/pkg/errors"
	"github.com/topics/models"
)

type DailyTrends struct {
	BasicCron
}

var trendModel = new(models.TrendModel)

const (
	locUS  = "TW"
	catAll = "all"
	langEn = "CN"
)

func (d *DailyTrends) Period() string {
	// return "@hourly"
	return "0 * * * *"
}

func (d *DailyTrends) Do() {
	log.Print("Daily trends start")
	//Enable debug to see request-response
	// gogtrends.Debug(true)
	ctx := context.Background()
	dailySearches, err := gogtrends.Daily(ctx, langEn, locUS)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Google trends - Daily"))
		return
	}
	for _, element := range dailySearches {
		count, err := strconv.Atoi(strings.Replace(*&element.FormattedTraffic, "K+", "000", 1))
		if err != nil {
			log.Fatal(errors.Wrap(err, "Item count replace K+ to 1000 fail"))
			return
		}
		trendModel.Store(count, *&element.Title.Query)
	}
	log.Println("Daily trends end")
}
