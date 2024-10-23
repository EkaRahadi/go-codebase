package main

import (
	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/httpserver"
	"github.com/EkaRahadi/go-codebase/internal/logger"
)

func main() {
	cfg := config.InitConfig()
	logger.SetZapLogger(cfg)
	defer logger.Log.Sync()

	httpserver.StartGinHttpServer(cfg)
}
