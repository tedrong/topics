package database

import (
	"time"

	"gorm.io/gorm"
)

type StockInfo struct {
	gorm.Model
	Symbol      string    `gorm:"type:text"`
	Name        string    `gorm:"type:text"`
	MarketType  string    `gorm:"type:text"`
	Industry    string    `gorm:"type:text"`
	ListingDate time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
}

// Market table have every day's TAIEX in Taiwan
type TAIEX struct {
	gorm.Model
	Date         time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	OpeningIndex float64   `gorm:"type:decimal(10,2);default:0"`
	ClosingIndex float64   `gorm:"type:decimal(10,2);default:0"`
	LowestIndex  float64   `gorm:"type:decimal(10,2);default:0"`
	HighestIndex float64   `gorm:"type:decimal(10,2);default:0"`
}

type DailyTrading struct {
	gorm.Model
	Symbol            string    `gorm:"type:text"`
	Date              time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	TradeVolume       int64     `gorm:"default:0"`
	TradeValue        int64     `gorm:"default:0"`
	OpeningPrice      float64   `gorm:"default:0"`
	HighestPrice      float64   `gorm:"default:0"`
	LowestPrice       float64   `gorm:"default:0"`
	ClosingPrice      float64   `gorm:"default:0"`
	Change            float64   `gorm:"default:0"`
	Transaction       int64     `gorm:"default:0"`
	DividendYield     float64   `gorm:"default:0"`
	DividendYear      string    `gorm:"type:text"`
	PERadio           float64   `gorm:"default:0"`
	PBRadio           float64   `gorm:"default:0"`
	FiscalYearQuarter string    `gorm:"type:text"`
}
type Highlight struct {
	gorm.Model
	Date        time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	TradeVolume float64   `gorm:"default:0"`
	TradeValue  float64   `gorm:"default:0"`
	Transaction float64   `gorm:"default:0"`
	TAIEX       float64   `gorm:"default:0"`
	Change      float64   `gorm:"default:0"`
}

type IndividualStockTrading struct {
	gorm.Model
	Date time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
}

// Trends table have every day's google search trends in Taiwan
type Trend struct {
	gorm.Model
	Date  time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	Rank  int       `gorm:"type:int"`
	Title string    `gorm:"type:text"`
	// Daily postgres.Jsonb `gorm:"type:jsonb;default:'{}'"`
}

// Content
type User struct {
	gorm.Model
	UUID     string
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}
