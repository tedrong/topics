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

func (c *Crawler) DailyTradingRatio(startDate time.Time) ([]*database.DailyTrading, error) {
	trades := []*database.DailyTrading{}
	// Make a date string with out hour/minut/second and convert back to time.Time
	strDate := (time.Now()).Format("2006-01-02")
	nowDate, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		log.Panic(errors.Wrap(err, "Time parsing fail"))
	}
	searchBtn, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main']//a[@class='button search']")
	category, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main']//select[@name='selectType']//option[contains(@value, 'ALL')]")

	for startDate.Before(nowDate) {
		// Input target year
		yearSelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='yy']//option[contains(@value, '%d')]", startDate.Year()))
		if err != nil {
			startDate = startDate.AddDate(0, 0, 1)
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

		// Input target day
		daySelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main']//div[@id='d1']//select[@name='dd']//option[contains(@value, '%d')]", int(startDate.Day())))
		if err != nil {
			return nil, err
		}
		err = daySelect.Click()
		if err != nil {
			return nil, err
		}

		err = category.Click()
		if err != nil {
			return nil, err
		}

		err = searchBtn.Click()
		if err != nil {
			return nil, err
		}

		time.Sleep(2 * time.Second)
		tableLength, err := (*c.WebDriver).FindElement(selenium.ByXPATH, "//div[@id='report-table_length']//select[@name='report-table_length']//option[contains(@value, -1)]")
		if err != nil {
			return nil, err
		}
		tableLength.Click()
		if err != nil {
			return nil, err
		}

		// Data table
		table, err := (*c.WebDriver).FindElement(selenium.ByID, "report-table")
		if err != nil {
			log.Panic(errors.Wrap(err, "FindElement: report-table"))
			continue
		}
		tableBody, err := table.FindElement(selenium.ByTagName, "tbody")
		if err != nil {
			log.Panic(errors.Wrap(err, "FindElement: table body"))
			continue
		}
		rows, _ := tableBody.FindElements(selenium.ByTagName, "tr")
		for _, row := range rows {
			columns, _ := row.FindElements(selenium.ByTagName, "td")
			symbol := getElenentText(&(columns[0]))
			trade, err := DailyTradingModel.GetBySymbolNDate(symbol, startDate)
			if err != nil {
				log.Panic(errors.Wrap(err, "Can't get record from database"))
				continue
			}
			if trade.Date == (time.Time{}) {
				trade.Symbol = symbol
				trade.Date = startDate
			}
			for idx, cell := range columns {
				switch idx {
				case 1:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					trade.DividendYield = value
				case 2:
					trade.DividendYear = getElenentText(&cell)
				case 3:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					trade.PERadio = value
				case 4:
					value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
					trade.PBRadio = value
				case 5:
					trade.FiscalYearQuarter = getElenentText(&cell)
				}
			}
			if trade.Date.Equal(startDate) || trade.Date.After(startDate) {
				trades = append(trades, trade)
			}
			// Break out if data structure length is larger then threshold
			if len(trades) >= 8192 {
				return trades, nil
			}
		}
		startDate = startDate.AddDate(0, 0, 1)
		// Sleep for a while, with random seed
		delay, err := strconv.Atoi(os.Getenv("CRAWLER_DELAY"))
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(common.RandInt(2, delay)) * time.Second)
	}
	return trades, nil
}
