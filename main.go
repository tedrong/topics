package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/topics/crawler"
	"github.com/topics/crontab"
	"github.com/topics/database"
	"github.com/topics/router"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Setting up log file rules
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./tmp/topics.log",
		MaxSize:    32, // megabytes
		MaxBackups: 2,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	})

	// Checking if environment is PRODUCTION, change gin to release mode
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Load the environment parameters file
	// Default file name is ".env", we can modify file name with Load() function. (e.g)godotenv.Load(./config/production.env)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	// Start PostgreSQL and Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	database.Init(1)
	seleniumService, c := crawler.StartWebInstance()
	crawler.StockInfo()

	TAIEXCron := crontab.TAIEX{BasicCron: crontab.BasicCron{}}
	DailyTrading := crontab.DailyTrading{BasicCron: crontab.BasicCron{}}
	Highlight := crontab.Highlight{BasicCron: crontab.BasicCron{}}
	trendsCron := crontab.DailyTrends{BasicCron: crontab.BasicCron{}}

	routine := cron.New()
	routine.AddFunc(TAIEXCron.Period(), TAIEXCron.Do)
	routine.AddFunc(trendsCron.Period(), trendsCron.Do)
	routine.AddFunc(DailyTrading.Period(), DailyTrading.Do)
	routine.AddFunc(Highlight.Period(), Highlight.Do)
	routine.Start()

	// Stop crawler and web driver
	defer (*c.WebDriver).Quit()
	defer seleniumService.Stop()

	router.Init()
}
