package earnings

import "time"

type Earnings struct {
	Symbol          string
	CompanyName     string
	Announcement    time.Time
	Period          string
	EPS             float64
	EPSForecast     float64
	Revenue         float64
	RevenueForecast float64
}
