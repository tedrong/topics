package models

import (
	"time"

	"github.com/topics/database"
	"github.com/topics/logging"
)

type TAIEXModel struct{}

func (m TAIEXModel) LatestDate() time.Time {
	zlog := logging.Get()
	db := database.GetPG(database.DBStock)
	row := database.TAIEX{}
	result := db.Last(&row)
	if result.Error != nil {
		date, err := time.Parse("2006-01-02", "1970-01-01")
		if err != nil {
			zlog.Panic().Err(err)
		}
		return date
	}
	return row.Date
}

func (m TAIEXModel) Store(markets []*database.TAIEX) {
	db := database.GetPG(database.DBStock)
	for _, element := range markets {
		if db.Model(&element).Where("date = ?", element.Date).Updates(&element).RowsAffected == 0 {
			db.Create(&element)
		}
	}
}
