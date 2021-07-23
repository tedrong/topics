package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Market table have every day's stock price index in Taiwan
type MarketIndex struct {
	gorm.Model
	Date         time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	OpeningPrice float64   `gorm:"type:decimal(10,2);default:0"`
	ClosingPrice float64   `gorm:"type:decimal(10,2);default:0"`
	MaxPrice     float64   `gorm:"type:decimal(10,2);default:0"`
	MinPrice     float64   `gorm:"type:decimal(10,2);default:0"`
}

func migrateMarketIndexTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&MarketIndex{})
	log.Print("The table 'market' migration is accomplished")
}
