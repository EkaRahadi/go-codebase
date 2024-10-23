package database

import (
	"context"
	"fmt"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitGorm(cfg *config.Config) *gorm.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Postgres.Host,
		dbCfg.Postgres.Username,
		dbCfg.Postgres.Password,
		dbCfg.Postgres.DbName,
		dbCfg.Postgres.Port,
		dbCfg.Sslmode,
	)

	var logCfg gormlogger.Config

	if cfg.App.Environment == constants.AppEnvironmentProduction {
		logCfg = gormlogger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
			LogLevel:                  gormlogger.Warn,
		}
	} else {
		logCfg = gormlogger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
			LogLevel:                  gormlogger.Info,
		}
	}

	newLogger := gormlogger.New(
		logger.Log,
		logCfg,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Log.Fatalw("error initializing database: ", err.Error())
	}

	// Add auto instrumentation
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		logger.Log.Fatalw("error instrumenting database: ", err.Error())
	}

	dbinstance, err := db.DB()
	if err != nil {
		logger.Log.Fatalw("error getting generic database instance: ", err.Error())
	}

	dbinstance.SetMaxIdleConns(dbCfg.MaxIdleConn)
	dbinstance.SetMaxOpenConns(dbCfg.MaxOpenConn)
	dbinstance.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db
}

// thin wrapper to enable transaction in repository
type GormWrapper struct {
	db *gorm.DB
}

func NewGormWrapper(db *gorm.DB) *GormWrapper {
	return &GormWrapper{
		db: db,
	}
}

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func (w *GormWrapper) Start(ctx context.Context) *gorm.DB {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return w.db.WithContext(ctx)
}
