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
	"github.com/topics/models"
)

var DailyTradingModel = new(models.DailyTrading)

func (c *Crawler) DailyTrading(stocks *[]database.StockInfo) ([]*database.DailyTrading, error) {
	trades := []*database.DailyTrading{}
	nowDate := time.Now()
	searchBtn, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main ajax']//a[@class='button search']")
	InputTextField, _ := (*c.WebDriver).FindElement(selenium.ByXPATH, "//form[@class='main ajax']//input[@name='stockNo']")

	for _, element := range *stocks {
		startDate := DailyTradingModel.LatestDate(element.Symbol)
		log.Printf("The latest date of stock - %s is %s", element.Symbol, startDate)
		for nowDate.After(startDate) {
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

			InputTextField.Clear()

			// Data table
			table, err := (*c.WebDriver).FindElement(selenium.ByID, "report-table")
			if err != nil {
				log.Fatal(errors.Wrap(err, "FindElement: report-table"))
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
							log.Fatal(errors.Wrap(err, "Time parsing fail"))
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
				// Take data if date if not the default value
				if trade.Date != (time.Time{}) {
					trades = append(trades, &trade)
				}
				// Break out if data structure length is larger then threshold
				if len(trades) >= 256 {
					return trades, nil
				}
			}
			startDate = startDate.AddDate(0, 1, 0)
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
