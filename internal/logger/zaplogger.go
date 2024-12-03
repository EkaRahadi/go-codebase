package logger

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	sugared *zap.SugaredLogger
}

func SetZapLogger(cfg *config.Config) {
	env := cfg.App.Environment
	logLevel := cfg.App.LogLevel // e.g., "debug", "info", "error"

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zapcore.InfoLevel // Default to Info level if invalid level is provided
	}

	var core zapcore.Core

	if env == constants.AppEnvironmentDevelopment {
		config := zap.NewDevelopmentEncoderConfig()
		config.TimeKey = "timestamp"
		encoder := zapcore.NewConsoleEncoder(config)

		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
	} else {
		udpAddr := fmt.Sprintf("%s:%d", cfg.Filebeat.Host, cfg.Filebeat.Port)

		conn, err := net.Dial("udp", udpAddr)
		if err != nil {
			log.Fatalf("failed to connect to filebeat via UDP: %v", err)
		}

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder := zapcore.NewJSONEncoder(config)
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		consoleWriter := zapcore.Lock(os.Stdout)

		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleWriter, level),
			zapcore.NewCore(encoder, zapcore.AddSync(conn), level), // UDP Core
		)
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap := &zapLogger{
		sugared: logger.Sugar().With("app_name", cfg.App.AppName),
	}

	SetLogger(zap)
}

func (zl *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	zl.sugared.Debugw(msg, keysAndValues...)
}

func (zl *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	zl.sugared.Infow(msg, keysAndValues...)
}

func (zl *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	zl.sugared.Warnw(msg, keysAndValues...)
}

func (zl *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	zl.sugared.Errorw(msg, keysAndValues...)
}

func (zl *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	zl.sugared.Fatalw(msg, keysAndValues...)
}

func (zl *zapLogger) Sync() error {
	return zl.sugared.Sync()
}

func (zl *zapLogger) getLogLevel(message string) (string, string) {
	prefixes := []string{"[warn]", "[error]", "[info]"}

	for _, prefix := range prefixes {
		if strings.Contains(message, prefix) {
			return prefix, message
		}
	}

	return "[info]", message
}

func (zl *zapLogger) Printf(format string, args ...interface{}) {
	level, msg := zl.getLogLevel(format)

	switch level {
	case constants.GormInfoLogLevel:
		zl.sugared.Infof(msg, args...)
	case constants.GormWarnLogLevel:
		zl.sugared.Warnw(msg, args...)
	case constants.GormErrorLogLevel:
		zl.sugared.Errorf(msg, args...)
	default:
		zl.sugared.Infof(msg, args...)
	}
}
