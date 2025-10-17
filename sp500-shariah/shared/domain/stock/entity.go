package stock

import "time"

type Stock struct {
	Symbol     string
	Name       string
	Date       time.Time
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Volume     int64
	Wvma       float64
	EarningsAt *time.Time
}
