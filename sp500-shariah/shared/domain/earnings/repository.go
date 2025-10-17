package earnings

type EarningsRepository interface {
	Save(e *Earnings) error
	FindBySymbol(symbol string) ([]*Earnings, error)
	FindLatest(symbol string) (*Earnings, error)
}
