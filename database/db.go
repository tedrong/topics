package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/topics/logging"
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
		{flag: DBStock, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_STOCK"))},
		{flag: DBTrend, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_TREND"))},
		{flag: DBContent, dsn: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME_CONTENT"))},
	}
	for _, element := range dbDSN {
		conn := ConnectDB(element.dsn)
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
func ConnectDB(dsn string) *gorm.DB {
	zlog := logging.Get()
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zlog.Panic().Err(err)
		return nil
	}
	return conn
}

func migration(flag DBFlag) {
	zlog := logging.Get()
	switch flag {
	case DBStock:
		// DBSet[flag].Debug().AutoMigrate(&StockInfo{}, &TAIEX{}, &DailyTrading{})
		DBSet[flag].AutoMigrate(&StockInfo{}, &TAIEX{}, &DailyTrading{}, &Highlight{})
		zlog.Debug().Msg("Table migrate successfully in DB:stock")
	case DBTrend:
		DBSet[flag].AutoMigrate(&Trend{})
		zlog.Debug().Msg("Table migrate successfully in DB:googleTrends")
	case DBContent:
		DBSet[flag].AutoMigrate(&User{})
		zlog.Debug().Msg("Table migrate successfully in DB:content")
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
