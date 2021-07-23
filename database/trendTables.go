package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Trends table have every day's google search trends in Taiwan
type Trend struct {
	gorm.Model
	Rank  int    `gorm:"type:int"`
	Title string `gorm:"type:text"`
	// Daily postgres.Jsonb `gorm:"type:jsonb;default:'{}'"`
}

func migrateTrendsTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&Trend{})
	log.Print("The table 'trends' migration is accomplished")
}
