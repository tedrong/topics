package models

import (
	"sort"
	"time"

	"github.com/topics/database"
	"github.com/topics/logging"
)

type DailyTrading struct{}

func (m DailyTrading) LatestDate(symbol string) time.Time {
	zlog := logging.Get()
	db := database.GetPG(database.DBStock)
	trades := []database.DailyTrading{}
	result := db.Where("symbol = ?", symbol).Where("trade_volume != ?", 0).Find(&trades)
	if result.RowsAffected == 0 || result.Error != nil {
		m := StockInfoModel{}
		info, err := m.GetBySymbol(symbol)
		if err != nil {
			zlog.Warn().Err(err)
		}
		return info.ListingDate
	}
	sort.Slice(trades, func(i, j int) bool {
		return trades[i].Date.After(trades[j].Date)
	})
	return trades[0].Date
}

func (m DailyTrading) LatestRatioDate() time.Time {
	zlog := logging.Get()
	db := database.GetPG(database.DBStock)
	trades := []database.DailyTrading{}
	result := db.Where("dividend_year = ?", "").Find(&trades)
	if result.RowsAffected == 0 || result.Error != nil {
		date, err := time.Parse("2006-01-02", "2005-09-02")
		if err != nil {
			zlog.Panic().Err(err)
		}
		return date
	}
	sort.Slice(trades, func(i, j int) bool {
		return trades[i].Date.Before(trades[j].Date)
	})
	return trades[0].Date
}

func (m DailyTrading) Store(markets []*database.DailyTrading) {
	db := database.GetPG(database.DBStock)
	for _, element := range markets {
		if db.Model(&element).Where("symbol = ?", element.Symbol).Where("date = ?", element.Date).Updates(&element).RowsAffected == 0 {
			db.Create(&element)
		}
	}
}

func (m DailyTrading) GetBySymbolNDate(symbol string, date time.Time) (*database.DailyTrading, error) {
	db := database.GetPG(database.DBStock)
	trade := database.DailyTrading{}
	result := db.Where("symbol = ?", symbol).Where("date = ?", date).Find(&trade)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trade, nil
}
