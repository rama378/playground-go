module github.com/rama378/playground-go/sp500-shariah/job-service

go 1.25.2

replace github.com/rama378/playground-go/sp500-shariah/shared => ../shared

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/lib/pq v1.10.9
	github.com/rama378/playground-go/sp500-shariah/shared v0.0.0-00010101000000-000000000000
)

require filippo.io/edwards25519 v1.1.0 // indirect
