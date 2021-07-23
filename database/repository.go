package database

import (
	"fmt"
	"log"
	"topics/config"

	"github.com/jinzhu/gorm"
)

type DatabaseConfig struct {
	flag DBFlag
	db   *gorm.DB
}

type DBFlag int

const (
	DBStock DBFlag = iota
	DBTrend
)

var DBSet [2]*DatabaseConfig

// ProvideDatabaseConfig declare dependency for wire
func ProvideDatabaseConfig(flag DBFlag) DatabaseConfig {
	return DatabaseConfig{flag: flag}
}

//Connect create the connection to postgresql and setting up gorm
func (p *DatabaseConfig) Connect() *DatabaseConfig {
	cfg := config.Get()
	var dbURI string
	switch p.flag {
	case DBStock:
		dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.DBStock.Host,
			cfg.DBStock.User,
			cfg.DBStock.Name,
			cfg.DBStock.Password)
	case DBTrend:
		dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.DBTrend.Host,
			cfg.DBTrend.User,
			cfg.DBTrend.Name,
			cfg.DBTrend.Password)
	}

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Print("Successfully connected to PostgreSQL")
	p.db = conn
	DBSet[p.flag] = p
	migration(p.flag)
	return p
}

func migration(flag DBFlag) {
	basicDTO := BasicDTO{}
	switch flag {
	case DBStock:
		migrateMarketIndexTable(basicDTO.Get(DBStock))
		log.Print("stockMarket database connection establashed")
	case DBTrend:
		migrateTrendsTable(basicDTO.Get(DBTrend))
		log.Print("googleTrends database connection establashed")
	}
}
