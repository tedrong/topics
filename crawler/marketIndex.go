package crawler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"topics/common"
	"topics/config"
	"topics/database"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
)

func (c *CrawlerEntry) MarketIndex(startDate time.Time) (*[]*database.MarketIndex, error) {
	cfg := config.Get()
	nowDate := time.Now()
	markets := []*database.MarketIndex{}
	searchBtn, _ := (*c.Crawler).FindElement(selenium.ByXPATH, "//form[@class='main']//a[@class='button search']")

	for nowDate.After(startDate) {
		//
		yearSelect, err := (*c.Crawler).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='yy']//option[contains(@value, '%d')]", startDate.Year()))
		if err != nil {
			startDate = startDate.AddDate(0, 1, 0)
			continue
		}
		monthSelect, _ := (*c.Crawler).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='mm']//option[contains(@value, '%d')]", int(startDate.Month())))

		err = yearSelect.Click()
		if err != nil {
			return nil, err
		}

		err = monthSelect.Click()
		if err != nil {
			return nil, err
		}

		err = searchBtn.Click()
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(common.RandInt(2, cfg.Crawler.Delay)) * time.Second)
		startDate = startDate.AddDate(0, 1, 0)

		// Data table
		table, err := (*c.Crawler).FindElement(selenium.ByID, "report-table")
		if err != nil {
			log.Fatal(errors.Wrap(err, "FindElement: report-table"))
		}
		rows, _ := table.FindElements(selenium.ByTagName, "tr")
		for _, row := range rows {
			market := database.MarketIndex{}
			columns, _ := row.FindElements(selenium.ByTagName, "td")
			for idx, cell := range columns {
				switch idx {
				case 0:
					// Split 110/01/01 to string slice
					seprateDate := strings.Split(getElenentText(&cell), "/")
					// Get 110 and add 1911 to make AD year
					year, _ := strconv.Atoi(seprateDate[0])
					seprateDate[0] = strconv.Itoa(year + 1911)
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
					market.Date = date
				case 1:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.OpeningPrice = price
				case 2:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.MaxPrice = price
				case 3:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.MinPrice = price
				case 4:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.ClosingPrice = price
				default:
				}
			}
			markets = append(markets, &market)
		}
		// Remove first element which contains default values
		markets = markets[1:]
	}
	return &markets, nil
}

func getElenentText(element *selenium.WebElement) string {
	value, err := (*element).Text()
	if err != nil {
		return ""
	}
	return value
}
