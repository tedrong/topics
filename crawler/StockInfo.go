package crawler

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"github.com/topics/database"
	"github.com/topics/models"
	"github.com/topics/sysexec"
)

var StockInfoModel = new(models.StockInfoModel)

func StockInfo() {
	// Check if there is a instance running, kill it
	if pid := sysexec.FindWebDriverPID(os.Getenv("CRAWLER_STOCKINFO_PORT")); pid != nil {
		sysexec.KillWebDriver(pid)
	}

	// Initialize
	crawlerEntry := CrawlerEntry{
		URL:             "https://isin.twse.com.tw/isin/class_main.jsp?owncode=&stockname=&isincode=&market=1&issuetype=1&industry_code=&Page=1&chklike=Y",
		SeleniumPath:    os.Getenv("SELENIUM"),
		GeckoDriverPath: os.Getenv("GECKO_DRIVER"),
	}

	port, err := strconv.Atoi(os.Getenv("CRAWLER_STOCKINFO_PORT"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "WebDriver can't get correct port number"))
	}
	crawlerEntry.Port = port

	// Driver instance startup
	webDriver, err := crawlerEntry.StartWebInstance()
	if err != nil {
		log.Fatal(errors.Wrap(err, "WebDriver instance startup fail"))
	}
	// New crawler with URL
	crawlerEntry.Crawler, err = crawlerEntry.Init()
	if err != nil {
		log.Fatal(errors.Wrap(err, "URL connection fail"))
	}

	infos := []*database.StockInfo{}
	tables, err := (*crawlerEntry.Crawler).FindElements(selenium.ByTagName, "table")
	if err != nil {
		log.Fatal(errors.Wrap(err, "FindElement: can't get any table"))
	}

	rows, _ := tables[1].FindElements(selenium.ByTagName, "tr")
	for idx, row := range rows {
		if idx != 0 {
			info := database.StockInfo{}
			columns, _ := row.FindElements(selenium.ByTagName, "td")
			for idx, cell := range columns {
				switch idx {
				case 2:
					info.Symbol = getElenentText(&cell)
				case 3:
					info.Name = getElenentText(&cell)
				case 4:
					info.MarketType = getElenentText(&cell)
				case 6:
					info.Industry = getElenentText(&cell)
				case 7:
					// Split 110/01/01 to string slice
					seprateDate := strings.Split(getElenentText(&cell), "/")
					// Make date string for lib parse
					strDate := ""
					for idx, element := range seprateDate {
						strDate += element
						if idx != len(seprateDate)-1 {
							strDate += "-"
						}
					}
					date, err := time.Parse("2006-01-02", strDate)
					if err != nil {
						log.Fatal(errors.Wrap(err, "Time parsing fail"))
					}
					info.ListingDate = date
				}
			}
			infos = append(infos, &info)
		}
	}
	StockInfoModel.Store(infos)

	// Stop crawler and web driver
	defer (*crawlerEntry.Crawler).Quit()
	defer webDriver.Stop()
}
