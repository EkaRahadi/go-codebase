package main

import (
	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/logger"
)

func main() {
	cfg := config.InitConfig()
	logger.SetZapLogger(cfg)
	defer logger.Log.Sync()

	logger.Log.Infow("test info log", "key", 123, "key1", "values")
	logger.Log.Errorw("test error log", "key", 123, "key1", "values")
}
