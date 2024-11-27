// pkg/log/logger.go
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global logger instance
var logger *zap.Logger

// InitLogger sets up the zap logger with proper configuration
func InitLogger(environment string) error {
	var config zap.Config

	if environment == "production" {
		// Production configuration: JSON format, info level and above
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// Development configuration: colored console output, debug level and above
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

// GetLogger returns the configured logger instance
func GetLogger() *zap.Logger {
	if logger == nil {
		// If logger hasn't been initialized, create a default development logger
		logger, _ = zap.NewDevelopment()
	}
	return logger
}
