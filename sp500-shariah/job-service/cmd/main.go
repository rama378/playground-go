package main

import (
	"log"

	"github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/db"
	postgres "github.com/rama378/playground-go/sp500-shariah/job-service/internal/infrastructure/mysql"
)

func main() {
	cfg := db.NewMySQLConfig()
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("❌ failed to connect to db: %v", err)
	}
	defer conn.Close()

	stockRepo := postgres.NewStockRepository(conn)
	log.Println("✅ StockRepository initialized:", stockRepo != nil)
}
