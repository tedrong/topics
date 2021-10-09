package crawler

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/tebeka/selenium"
	"github.com/topics/common"
	"github.com/topics/database"
	"github.com/topics/logging"
)

func (c *Crawler) TAIEX(startDate time.Time) ([]*database.TAIEX, error) {
	markets := []*database.TAIEX{}
	// Make a date string with out hour/minut/second and convert back to time.Time
	strDate := (time.Now()).Format("2006-01-02")
	nowDate, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		c.LogJob(logging.Get().Panic(), CR_TAIEX).Err(err)
	}
	searchBtn, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main']//a[@class='button search']")

	for startDate.Before(nowDate) {
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
			c.LogJob(logging.Get().Panic(), CR_TAIEX).Err(err)
			continue
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
						c.LogJob(logging.Get().Panic(), CR_TAIEX).Err(err)
						break
					}
					market.Date = date
				case 1:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.OpeningIndex = value
				case 2:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.HighestIndex = value
				case 3:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.LowestIndex = value
				case 4:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					market.ClosingIndex = value
				default:
				}
			}
			// Take data if date if not the default value
			if (market.Date.Equal(startDate) || market.Date.After(startDate)) && market.Date != (time.Time{}) {
				markets = append(markets, &market)
			}
		}
		if startDate.Day() == (now.With(startDate).EndOfMonth()).Day() {
			startDate = startDate.AddDate(0, 0, 1)
		} else {
			startDate = startDate.AddDate(0, 1, 0)
		}
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
