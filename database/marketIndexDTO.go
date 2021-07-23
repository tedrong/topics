package database

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type MarketIndexDTO struct {
	BasicDTO
	MarketIndex MarketIndex
}

func (m *MarketIndexDTO) Insert() {
	db := m.Get(DBStock)
	if db.Model(&m.MarketIndex).Where("date = ?", m.MarketIndex.Date).Updates(&m.MarketIndex).RowsAffected == 0 {
		db.Create(&m.MarketIndex)
	}
}

func (m *MarketIndexDTO) LatestDate() time.Time {
	db := m.Get(DBStock)
	result := db.Last(&m.MarketIndex)
	if result.Error != nil {
		date, err := time.Parse("2006-01-02", "1970-01-01")
		if err != nil {
			log.Fatal(errors.Wrap(err, "Time parsing fail"))
		}
		return date
	}
	return m.MarketIndex.Date
}
