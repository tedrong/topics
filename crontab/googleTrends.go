package crontab

import (
	"context"
	"log"
	"strconv"
	"strings"
	"topics/database"

	"github.com/groovili/gogtrends"
	"github.com/pkg/errors"
)

type DailyTrends struct {
	BasicCron
}

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
		dto := database.TrendDTO{
			BasicDTO: database.BasicDTO{},
			Trend: database.Trend{
				Rank:  count,
				Title: *&element.Title.Query,
			}}
		record, err := dto.FetchByTitle()
		if err != nil {
			log.Println("New trend record, insert to database")
			dto.Insert()
		}

		if record != nil && count > record.Rank {
			dto.Trend.Model = record.Model
			err := dto.Update()
			if err != nil {
				log.Printf("Can not update title(%s), err: %s", dto.Trend.Title, err)
			}
		}
	}
	log.Println("Daily trends end")
}
