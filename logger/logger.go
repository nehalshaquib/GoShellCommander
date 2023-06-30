package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() (*zap.SugaredLogger, error) {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout","./logs.log"],
		"errorOutputPaths": ["stderr","./logs.log"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return logger.Sugar(), nil
}
