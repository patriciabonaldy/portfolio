package fintual

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func randate() time.Time {
	min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2021, 9, 29, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (p Portfolio) AddStock(stock Stock)  {
	mapStocks, ok := p.stocks[stock.date]
	if !ok {
		mapStocks:= make(map[int64][]Stock)
		mapStocks[stock.ID] = []Stock{stock}
		p.stocks[stock.date] = mapStocks
		return
	}

	stocks := mapStocks[stock.ID]
	stocks = append(stocks, stock)
	mapStocks[stock.ID] = stocks
	p.stocks[stock.date] = mapStocks
}

func (s Stock) Price(date time.Time) (float64, error) {
	historyStock, ok := HistoryStock[date]
	if !ok {
		return 0, fmt.Errorf("not exists data for this date")
	}

	return historyStock.Value, nil
}

func (p Portfolio) AnnualizedProfit() (float64, error) {
	var totalAnnualized float64
	if len(p.stocks) == 0 {
		return 0, fmt.Errorf("error: portfolio does not have stocks")
	}

	for date, stocks := range p.stocks {
		today := time.Now()
		days := 365/(today.Sub(date).Hours() / 24)-1
		for _, _stocks := range stocks {
			cost, err := p.calculateCost(date, _stocks)
			if err != nil {
				return 0, err
			}
			proceeds, err := p.calculateProceeds(_stocks)
			if err != nil {
				return 0, err
			}
			annualized := math.Pow(proceeds/cost, days)
			totalAnnualized += annualized
		}
	}

	return totalAnnualized*100, nil
}

func (p Portfolio) Profit(start, end time.Time) (float64, error) {
	var totalCost,totalProceeds float64
	if len(p.stocks) == 0 {
		return 0, fmt.Errorf("portfolio does not have stocks")
	}

	for date, stocks := range p.stocks {
		// verify that the param date is in range of date
		if inTimeSpan(start, end, date) {
			for _, _stocks := range stocks {
				cost, err := p.calculateCost(date, _stocks)
				if err != nil {
					return 0, err
				}

				totalCost += cost
				proceeds, err := p.calculateProceeds(_stocks)
				if err != nil {
					return 0, err
				}

				totalProceeds += proceeds
			}
		}
	}

	profit := totalProceeds - totalCost

	return profit, nil
}

func (p Portfolio) calculateCost(date time.Time, stocks []Stock) (float64, error) {
	if len(stocks) == 0 {
		return 0, nil
	}

	price, err := stocks[0].Price(date)
	if err != nil {
		return 0, nil
	}

	return float64(len(stocks)) * price, nil
}

func (p Portfolio) calculateProceeds(stocks []Stock) (float64, error) {
	if len(stocks) == 0 {
		return 0, nil
	}
	interestRate:= (float64(rand.Intn(3)) / float64(2)*1.05)+ 1

	return float64(len(stocks)) * stocks[0].Value * interestRate, nil
}
