package fintual

import "time"

var HistoryStock = make(map[time.Time]Stock)

type Portfolio struct {
	stocks map[time.Time]map[int64][]Stock
}

type Stock struct {
	ID    int64
	Value float64
	date  time.Time
}
