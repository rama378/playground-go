package application

import (
	"fmt"
	"time"

	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/yahoo"
	"github.com/rama378/playground-go/sp500-shariah/shared/domain/stock"
)

type StockService struct {
	yahooClient *yahoo.YahooClient
	repo        stock.StockRepository
}

func NewStockService(client *yahoo.YahooClient, repo stock.StockRepository) *StockService {
	return &StockService{yahooClient: client, repo: repo}
}

func (s *StockService) FetchAndSaveStock(symbol string) error {
	data, err := s.yahooClient.FetchStockData(symbol)
	if err != nil {
		return fmt.Errorf("fetch failed: %w", err)
	}

	if len(data.Chart.Result) == 0 {
		return fmt.Errorf("no data returned for symbol %s", symbol)
	}

	result := data.Chart.Result[0]
	if len(result.Indicators.Quote) == 0 {
		return fmt.Errorf("no quote data for symbol %s", symbol)
	}

	for i, t := range result.Timestamp {
		close := result.Indicators.Quote[0].Close[i]
		stock := stock.Stock{
			Symbol: result.Meta.Symbol,
			Date:   time.Unix(0, t),
			Close:  close,
		}
		if err := s.repo.Save(&stock); err != nil {
			return fmt.Errorf("save failed: %w", err)
		}
	}

	return nil
}
