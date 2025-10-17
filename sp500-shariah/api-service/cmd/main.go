package main

import (
	"fmt"

	"github.com/rama378/playground-go/sp500-shariah/shared/config"
	"github.com/rama378/playground-go/sp500-shariah/shared/logger"
)

func main() {
	cfg := config.Load("configs/config.yaml")
	logger.Info(fmt.Sprintf("%s started in %s mode", cfg.AppName, cfg.Env))
}
