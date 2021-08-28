package crawler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"github.com/topics/common"
	"github.com/topics/database"
)

func (c *Crawler) TAIEX(startDate time.Time) ([]*database.TAIEX, error) {
	markets := []*database.TAIEX{}
	nowDate := time.Now()
	searchBtn, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main']//a[@class='button search']")

	for nowDate.After(startDate) {
		// Input target year
		yearSelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='yy']//option[contains(@value, '%d')]", startDate.Year()))
		if err != nil {
			startDate = startDate.AddDate(0, 1, 0)
			continue
		}
		// Input target month
		monthSelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='mm']//option[contains(@value, '%d')]", int(startDate.Month())))
		if err != nil {
			return nil, err
		}

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
		// Data table
		table, err := (*c.WebDriver).FindElement(selenium.ByID, "report-table")
		if err != nil {
			log.Fatal(errors.Wrap(err, "FindElement: report-table"))
		}
		rows, _ := table.FindElements(selenium.ByTagName, "tr")
		for _, row := range rows {
			market := database.TAIEX{}
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
					market.OpeningIndex = price
				case 2:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.HighestIndex = price
				case 3:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.LowestIndex = price
				case 4:
					price, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.ClosingIndex = price
				default:
				}
			}
			// Take data if date if not the default value
			if market.Date != (time.Time{}) {
				markets = append(markets, &market)
			}
		}
		startDate = startDate.AddDate(0, 1, 0)
		// Break out loop if data structure length is larger then threshold
		if len(markets) >= 256 {
			break
		}
		// Sleep for a while, with random seed
		delay, err := strconv.Atoi(os.Getenv("CRAWLER_DELAY"))
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(common.RandInt(2, delay)) * time.Second)
	}
	return markets, nil
}
