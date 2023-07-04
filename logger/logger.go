package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

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

	// Create log file with current timestamp
	logPath := fmt.Sprintf("%s/%s.log", dir, currentTime.Format("2023-07-03_15-04-05"))

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
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return logger.Sugar(), nil
}
