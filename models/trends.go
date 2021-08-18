package models

import (
	"log"

	"github.com/topics/database"
)

type TrendModel struct{}

func (m TrendModel) Store(rank int, title string) {
	db := database.GetPG(database.DBTrend)
	row := database.Trend{
		Rank:  rank,
		Title: title,
	}
	if db.Model(&row).Where("title = ?", title).Updates(row).RowsAffected == 0 {
		db.Create(&row)
	} else {
		log.Printf("Title(%s) updated with Rank(%d)", title, rank)
	}
}
