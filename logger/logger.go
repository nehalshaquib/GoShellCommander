package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes a new logger instance and returns it.
// It creates a new log file in the "./logs" directory with the current timestamp as its name.
// If the "./logs" directory does not exist, it creates it.
// The logger is configured to write both to the standard output and the log file.
// It uses the "json" encoding and the production encoder.
// The log level is set to "info", meaning that all "info", "warn", and "error" messages will be logged.
// Debug messages will be ignored.
// The logger uses a sampling strategy to keep the logging cost constant.
// The initial number of messages to log is 100, and thereafter, every 100th message is logged.
func InitLogger() (*zap.SugaredLogger, error) {
	currentTime := time.Now()
	dir := "./logs"

	// Create directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return nil, errors.New("failed to create directory:" + err.Error())
		}
	}

	// Create the log file path with the current timestamp
	logPath := fmt.Sprintf("%s/%s.log", dir, currentTime.Format("D-2006-01-02_T-15-04-05"))

	// Configure the logger
	cfg := zap.Config{
		OutputPaths: []string{
			"stdout",
			logPath,
		},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:      "json",
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Build the logger from the configuration
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// Ensure that all pending log entries are written
	defer logger.Sync()

	// Return the logger
	return logger.Sugar(), nil
}
