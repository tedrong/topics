package crontab

import (
	"context"
	"strconv"
	"strings"

	"github.com/groovili/gogtrends"
	"github.com/topics/logging"
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
	d.LogJob(logging.Get().Info(), CJ_Trends).Msg("Daily trends start")
	//Enable debug to see request-response
	// gogtrends.Debug(true)
	ctx := context.Background()
	dailySearches, err := gogtrends.Daily(ctx, langEn, locUS)
	if err != nil {
		d.LogJob(logging.Get().Warn(), CJ_Trends).Err(err)
		return
	}
	for _, element := range dailySearches {
		count, err := strconv.Atoi(strings.Replace(element.FormattedTraffic, "K+", "000", 1))
		if err != nil {
			d.LogJob(logging.Get().Warn(), CJ_Trends).Err(err)
			return
		}
		trendModel.Store(count, element.Title.Query)
	}
	d.LogJob(logging.Get().Info(), CJ_Trends).Msg("Daily trends end")
}
