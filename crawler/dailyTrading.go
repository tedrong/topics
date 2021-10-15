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
	"github.com/topics/models"
)

var DailyTradingModel = new(models.DailyTrading)

func (c *Crawler) DailyTrading(stocks *[]database.StockInfo) ([]*database.DailyTrading, error) {
	// Declare the row data in database
	trades := []*database.DailyTrading{}
	// Make a date string with out hour/minut/second and convert back to time.Time
	strDate := (time.Now()).Format("2006-01-02")
	nowDate, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		c.LogJob(logging.Get().Panic(), CR_DailyTrading).Err(err)
	}
	// Locate the search button and stock symbol input field
	searchBtn, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main ajax']//a[@class='button search']")
	InputTextField, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main ajax']//input[@name='stockNo']")

	// Loop for stocks
	for _, element := range *stocks {
		// Get latest stock information from database
		startDate := DailyTradingModel.LatestDate(element.Symbol)
		// Skip stock if it's already has newest information
		if startDate.Equal(nowDate.AddDate(0, 0, -1)) {
			continue
		}
		c.LogJob(logging.Get().Debug(), CR_DailyTrading).Msg(fmt.Sprintf("The latest date of stock - %s is %s", element.Symbol, startDate))
		// Loop for dates
		for startDate.Before(nowDate) {
			// Input target year
			yearSelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main ajax']//div[@id='d1']//select[@name='yy']//option[contains(@value, '%d')]", startDate.Year()))
			if err != nil {
				startDate = startDate.AddDate(0, 1, 0)
				continue
			}
			// Input target month
			monthSelect, err := (*c.WebDriver).FindElement(selenium.ByXPATH, fmt.Sprintf("//form[@class='main ajax']//div[@id='d1']//select[@name='mm']//option[contains(@value, '%d')]", int(startDate.Month())))
			if err != nil {
				return nil, err
			}

			// Input stock symbol to the text field
			InputTextField.SendKeys(element.Symbol)

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

			time.Sleep(2 * time.Second)
			InputTextField.Clear()

			// Data table
			table, err := (*c.WebDriver).FindElement(selenium.ByID, "report-table")
			if err != nil {
				c.LogJob(logging.Get().Panic(), CR_DailyTrading).Err(err)
				continue
			}
			rows, _ := table.FindElements(selenium.ByTagName, "tr")
			for _, row := range rows {
				trade := database.DailyTrading{}
				trade.Symbol = element.Symbol
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
							c.LogJob(logging.Get().Panic(), CR_DailyTrading).Err(err)
							break
						}
						trade.Date = date
					case 1:
						value, _ := strconv.ParseInt(strings.ReplaceAll(getElenentText(&cell), ",", ""), 10, 64)
						trade.TradeVolume = value
					case 2:
						value, _ := strconv.ParseInt(strings.ReplaceAll(getElenentText(&cell), ",", ""), 10, 64)
						trade.TradeValue = value
					case 3:
						value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
						trade.OpeningPrice = value
					case 4:
						value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
						trade.HighestPrice = value
					case 5:
						value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
						trade.LowestPrice = value
					case 6:
						value, _ := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
						trade.ClosingPrice = value
					case 7:
						value, err := strconv.ParseFloat(strings.ReplaceAll(getElenentText(&cell), ",", ""), 64)
						if err != nil {
							value = 0
						}
						trade.Change = value
					case 8:
						value, _ := strconv.ParseInt(strings.ReplaceAll(getElenentText(&cell), ",", ""), 10, 64)
						trade.Transaction = value
					}
				}
				// Take data if there are new result from table, and data time shouldn't be default value
				if (trade.Date.Equal(startDate) || trade.Date.After(startDate)) && trade.Date != (time.Time{}) {
					trades = append(trades, &trade)
				}
			}
			if startDate.Day() == (now.With(startDate).EndOfMonth()).Day() {
				startDate = startDate.AddDate(0, 0, 1)
			} else {
				startDate = startDate.AddDate(0, 1, 0)
			}
			// Break out if data structure length is larger then threshold
			if len(trades) >= 256 {
				return trades, nil
			}
			// Sleep for a while, with random seed
			delay, err := strconv.Atoi(os.Getenv("CRAWLER_DELAY"))
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Duration(common.RandInt(2, delay)) * time.Second)
		}
	}
	return trades, nil
}
