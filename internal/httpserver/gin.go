package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/handler/ginhandler"
	"github.com/EkaRahadi/go-codebase/internal/httpserver/ginroutes"
	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/EkaRahadi/go-codebase/internal/telemetry"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func StartGinHttpServer(cfg *config.Config) {

	// Telemetry Setup
	otel := telemetry.NewTelemetry()
	otel.InitGlobalProviderOpenTelemetry(cfg.Otlp.OtelExporterOtlpMetricsEndpoint, cfg.App.AppName)

	// database
	db := database.InitGorm(cfg)
	// redis := database.InitRedis(cfg)

	// dependencies
	gormWrapper := database.NewGormWrapper(db)
	transactor := database.NewTransactor(db)
	// redisWrapper := database.NewRedisWrapper(redis)
	vldtr := utils.NewCustomValidator()
	jwtUtil := utils.NewJWTUtil(cfg)
	// authUtil := utils.NewAuthUtil(cfg)

	appHandler := ginhandler.NewAppHandler()

	if cfg.App.Environment == constants.AppEnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.ContextWithFallback = true // enables .Done(), .Err(), and .Value()
	r.MaxMultipartMemory = cfg.HttpServer.MaxUploadFileSize

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*", "http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Content-Type,access-control-allow-origin, access-control-allow-headers", "Authorization"}
	r.Use(cors.New(config))

	// registering middlewares
	middlewares := []gin.HandlerFunc{
		// middleware.RequestId(),
		middleware.Logger(),
		middleware.ErrorHandler(),
		gin.Recovery(),
	}
	// Setup gin auto instrumentation
	r.Use(otelgin.Middleware(cfg.App.AppName))

	r.Use(middlewares...)
	r.NoRoute(appHandler.RouteNotFound)

	/*
		All Routes here
	*/
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Welcome to %s BE Server", cfg.App.AppName)})
	})
	// ginroutes.RealRoutes(r, gormWrapper, vldtr)

	if cfg.App.Environment == constants.AppEnvironmentDevelopment {
		ginroutes.RegisterExampleRoutes(r, gormWrapper, transactor, vldtr)
		ginroutes.RegisterTokenRoutes(r, gormWrapper, vldtr, jwtUtil)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	go func() {
		logger.Log.Infow("running server on port :", "listen_port", cfg.HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatalw("error while server listen and serve", "error: ", err)
			}
		}
		logger.Log.Infow("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(cfg.HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	otel.Shutdown(ctx)
	defer cancel()

	logger.Log.Infow("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatalw("error shutting down server: ", "error: ", err)
	}

	logger.Log.Infow("http server is shutting down gracefully")
}
