package main

import (
	"github.com/rama378/playground-go/sp500-shariah/shared/config"
	"github.com/rama378/playground-go/sp500-shariah/shared/logger"
)

func main() {
	cfg := config.Load("configs/config.yaml")

	logger.Info("%s started in %s mode", cfg.AppName, cfg.Env)
	logger.Warm("API routes not yet registered")
}
