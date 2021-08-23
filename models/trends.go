package models

import (
	"log"
	"time"

	"github.com/topics/database"
)

type TrendModel struct{}

func (m TrendModel) Store(rank int, title string) {
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
		log.Printf("Title(%s) updated with Rank(%d)", title, rank)
	}
}
