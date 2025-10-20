package db

import (
	"database/sql"
	"log"
)

func AutoMigrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS stocks (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		symbol VARCHAR(20) NOT NULL,
		name VARCHAR(20) NOT NULL,
		date DATE NOT NULL,
		open DECIMAL(10,2),
		high DECIMAL(10,2),
		low DECIMAL(10,2),
		close DECIMAL(10,2),
		volume BIGINT,
		wvma DECIMAL(10,2),
		earnings_at DATETIME NULL,
		UNIQUE KEY unique_symbol_date (symbol, date)
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ stocks table initialized")
	return nil
}
