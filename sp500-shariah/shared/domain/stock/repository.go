package stock

import "time"

type StockRepository interface {
	Save(stock *Stock) error
	FindBySymbol(symbol string) ([]*Stock, error)
	FindLatest(symbol string) (*Stock, error)
	FindByDateRange(symbol string, from, to time.Time) ([]*Stock, error)
}
