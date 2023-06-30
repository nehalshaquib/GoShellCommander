package config

import (
	"github.com/nehalshaquib/GoShellCommander/logger"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

func Configure() error {
	log, err := logger.InitLogger()
	if err != nil {
		return err
	}
	Logger = log

	return nil
}
