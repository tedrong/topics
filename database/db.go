package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
type DBFlag int

const (
	DBStock DBFlag = iota
	DBTrend
	DBContent
)

type DB struct {
	flag DBFlag
	dsn  string
}

var DBSet [3]*gorm.DB
var RedisClient *redis.Client

// Init ...
func Init(selectDB ...int) {
	dbDSN := []DB{
		{flag: DBStock, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_STOCK"))},
		{flag: DBTrend, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_TREND"))},
		{flag: DBContent, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_CONTENT"))},
	}
	for _, element := range dbDSN {
		conn, err := ConnectDB(element.dsn)
		if err != nil {
			log.Panic(err)
		}
		DBSet[element.flag] = conn
		migration(element.flag)
	}

	//InitRedis ...
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})
}

// Connect create the connection to postgresql and setting up gorm
func ConnectDB(dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func migration(flag DBFlag) {
	switch flag {
	case DBStock:
		DBSet[flag].Debug().AutoMigrate(&StockInfo{}, &TAIEX{}, &DailyTrading{})
		log.Print("Table migrate successfully in DB:stock")
	case DBTrend:
		DBSet[flag].Debug().AutoMigrate(&Trend{})
		log.Print("Table migrate successfully in DB:googleTrends")
	case DBContent:
		DBSet[flag].Debug().AutoMigrate(&User{})
		log.Print("Table migrate successfully in DB:content")
	}
}

// GetPG ...
func GetPG(flag DBFlag) *gorm.DB {
	return DBSet[flag]
}

//GetRedis ...
func GetRedis() *redis.Client {
	return RedisClient
}
