package database

import (
	"time"

	"gorm.io/gorm"
)

// Market table have every day's stock price index in Taiwan
type TAIEX struct {
	gorm.Model
	Date         time.Time `gorm:"type:timestamp with time zone;default:'1970-01-01 0:00AM'"`
	OpeningIndex float64   `gorm:"type:decimal(10,2);default:0"`
	ClosingIndex float64   `gorm:"type:decimal(10,2);default:0"`
	LowestIndex  float64   `gorm:"type:decimal(10,2);default:0"`
	HighestIndex float64   `gorm:"type:decimal(10,2);default:0"`
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
