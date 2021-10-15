package crawler

import (
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/topics/database"
	"github.com/topics/logging"
	"github.com/topics/models"
)

var StockInfoModel = new(models.StockInfoModel)

func StockInfo() {
	crawler := Get()
	crawler.Mutex.Lock()
	rows := StockInfoModel.All()
	if len(*rows) == 0 {
		crawler.URL = "https://isin.twse.com.tw/isin/class_main.jsp?owncode=&stockname=&isincode=&market=1&issuetype=1&industry_code=&Page=1&chklike=Y"
		crawler.GOTO()

		infos := []*database.StockInfo{}
		tables, err := (*crawler.WebDriver).FindElements(selenium.ByTagName, "table")
		if err != nil {
			crawler.LogJob(logging.Get().Panic(), CR_StockInfo).Err(err)
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
							crawler.LogJob(logging.Get().Panic(), CR_StockInfo).Err(err)
						}
						info.ListingDate = date
					}
				}
				infos = append(infos, &info)
			}
		}
		StockInfoModel.Store(infos)
	}
	crawler.Mutex.Unlock()
}
