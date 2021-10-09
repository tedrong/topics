package models

import (
	"fmt"
	"time"

	"github.com/topics/database"
	"github.com/topics/logging"
)

type TrendModel struct{}

func (m TrendModel) Store(rank int, title string) {
	zlog := logging.Get()
	db := database.GetPG(database.DBTrend)
	date := time.Now()
	row := database.Trend{
		Date:  time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local),
		Rank:  rank,
		Title: title,
	}
	if db.Model(&row).Where("title = ?", title).Where("date = ?", row.Date).Updates(row).RowsAffected == 0 {
		db.Create(&row)
	} else {
		zlog.Info().Msg(fmt.Sprintf("Title(%s) updated with Rank(%d)", title, rank))
	}
}
