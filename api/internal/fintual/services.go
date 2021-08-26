package fintual

import (
	"math/rand"
	"time"
)

// StockService interface .
type StockService interface {
	Price(date time.Time) (float64, error)
}

// PortfolioService interface
type PortfolioService interface {
	AddStock(stock Stock)
	AnnualizedProfit() (float64, error)
	Profit(start, end time.Time) (float64, error)
}

func NewStock() Stock {
	const minRand = 100
	const maxRand = 1000
	date:= randate()
	value := float64(rand.Intn((maxRand-minRand)/0.05))*0.05 + minRand
	stock := Stock{
		ID:    int64(rand.Intn(200)),
		Value: value,
		date:  date,
	}
	HistoryStock[date] = stock
	return stock
}

func NewPortfolio() PortfolioService {
	return &Portfolio{stocks: make(map[time.Time]map[int64][]Stock, 0)}
}