package postgres

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/rama378/playground-go/sp500-shariah/shared/domain/stock"
)

type StockRepository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Save(s *stock.Stock) error {
	query := `
        INSERT INTO stocks (symbol, name, date, open, high, low, close, volume, wvma, earnings_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
        ON CONFLICT (symbol, date) DO UPDATE 
        SET open=$4, high=$5, low=$6, close=$7, volume=$8, wvma=$9, earnings_at=$10;
    `

	_, err := r.db.Exec(query, s.Symbol, s.Name, s.Date, s.Open, s.High, s.Low, s.Close, s.Volume, s.Wvma, s.EarningsAt)

	return err
}

func (r *StockRepository) FindBySymbol(symbol string) ([]*stock.Stock, error) {
	query := `SELECT symbol, name, date, open, high, low, close, volume, wvma, earnings_at FROM stocks WHERE symbol=$1 ORDER BY date`
	rows, err := r.db.Query(query, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*stock.Stock
	for rows.Next() {
		var s stock.Stock
		var earningsAt sql.NullTime
		if err := rows.Scan(
			&s.Symbol, &s.Name, &s.Date, &s.Open, &s.High, &s.Low, &s.Close, &s.Volume, &s.Wvma, &earningsAt,
		); err != nil {
			return nil, err
		}

		if earningsAt.Valid {
			s.EarningsAt = &earningsAt.Time
		}
	}
	return result, nil
}

func (r *StockRepository) FindLatest(symbol string) (*stock.Stock, error) {
	query := `SELECT symbol, name, date, open, high, low, close, volume, wvma, earnings_at FROM stocks WHERE symbol=$1 ORDER BY date DESC LIMIT 1`
	row := r.db.QueryRow(query, symbol)

	var s stock.Stock
	var earningsAt sql.NullTime
	if err := row.Scan(&s.Symbol, &s.Name, &s.Date, &s.Open, &s.High, &s.Low, &s.Close, &s.Volume, &s.Wvma, &earningsAt); err != nil {
		return nil, err
	}
	if earningsAt.Valid {
		s.EarningsAt = &earningsAt.Time
	}

	return &s, nil
}

func (r *StockRepository) FindByDateRange(symbol string, from, to time.Time) ([]*stock.Stock, error) {
	query := `SELECT symbol, name, date, open, high, low, close, volume, wvma, earnings_at FROM stocks WHERE symbol=$1 AND date BETWEEN $2 AND $3 ORDER BY date`
	rows, err := r.db.Query(query, symbol, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*stock.Stock
	for rows.Next() {
		var s stock.Stock
		var earningsAt sql.NullTime
		if err := rows.Scan(
			&s.Symbol, &s.Name, &s.Date, &s.Open, &s.High, &s.Low, &s.Close, &s.Volume, &s.Wvma, &earningsAt,
		); err != nil {
			return nil, err
		}

		if earningsAt.Valid {
			s.EarningsAt = &earningsAt.Time
		}
		result = append(result, &s)
	}

	return result, nil
}
