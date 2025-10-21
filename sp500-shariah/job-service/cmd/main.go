package main

import (
	"log"

	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/db"
	mysql "github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/mysql"
	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/interfaces/job"
)

func main() {
	cfg := db.NewMySQLConfig()
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("❌ failed to connect to db: %v", err)
	}
	defer conn.Close()

	if err := db.AutoMigrate(conn); err != nil {
		log.Fatalf("❌ migration failed: %v", err)
	}

	stockRepo := mysql.NewStockRepository(conn)
	job.RunYahooJob(stockRepo)
	log.Println("✅ StockRepository initialized:", stockRepo != nil)
}
