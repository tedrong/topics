package models

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/topics/database"
)

type DailyTradingModel struct{}

func (m DailyTradingModel) LatestDate() time.Time {
	db := database.GetPG(database.DBStock)
	row := database.DailyTrading{}
	result := db.Last(&row)
	if result.Error != nil {
		date, err := time.Parse("2006-01-02", "1970-01-01")
		if err != nil {
			log.Fatal(errors.Wrap(err, "Time parsing fail"))
		}
		return date
	}
	return row.Date
}

func (m DailyTradingModel) Store(markets []*database.DailyTrading) {
	db := database.GetPG(database.DBStock)
	for _, element := range markets {
		if db.Model(&element).Where("date = ?", element.Date).Updates(&element).RowsAffected == 0 {
			db.Create(&element)
		}
	}
}
