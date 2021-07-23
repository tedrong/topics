package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/natefinch/lumberjack.v2"

	"topics/config"
	"topics/crontab"
	"topics/database"
)

const (
	configFileName = "topics.yml"
)

func main() {
	// Get the configurations from config file
	cfg := config.Load(configFileName)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   cfg.Log.Path + cfg.Log.Name,
		MaxSize:    32, // megabytes
		MaxBackups: 2,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	})

	// Databases initialize and connect
	stockDB := InitDB(database.DBStock)
	trendDB := InitDB(database.DBTrend)
	stockDB.Connect()
	trendDB.Connect()

	marketIndexCron := crontab.MarketIndex{BasicCron: crontab.BasicCron{}}
	go marketIndexCron.Do()

	// trendsCron := crontab.DailyTrends{BasicCron: crontab.BasicCron{}}
	// trendsCron.Do()

	// routine := cron.New()
	// routine.AddFunc(trendsCron.Period(), trendsCron.Do)
	// routine.Start()

	basicDTO := database.BasicDTO{}
	productAPI := InitProductAPI(basicDTO.Get(database.DBStock))
	r := gin.Default()
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
