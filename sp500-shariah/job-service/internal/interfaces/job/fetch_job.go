package job

import (
	"fmt"

	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/application"
	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/yahoo"
	"github.com/rama378/playground-go/sp500-shariah/shared/domain/stock"
)

func RunYahooJob(repo stock.StockRepository) {
	yahooClient := yahoo.NewYahooClient()
	service := application.NewStockService(yahooClient, repo)

	symbols := []string{"AAPL", "GOOG", "MSFT"}
	for _, sym := range symbols {
		fmt.Printf("Fetching %s... \n", sym)
		if err := service.FetchAndSaveStock(sym); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
