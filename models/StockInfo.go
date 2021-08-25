package models

import "github.com/topics/database"

type StockInfoModel struct{}

func (m StockInfoModel) Store(info []*database.StockInfo) {
	db := database.GetPG(database.DBStock)
	for _, element := range info {
		if db.Model(&element).Where("symbol = ?", element.Symbol).Updates(&element).RowsAffected == 0 {
			db.Create(&element)
		}
	}
}
