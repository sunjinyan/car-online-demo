package server

import (
	"go.uber.org/zap"
)

func CreateLogger() (*zap.Logger,error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = ""
	return  config.Build()
}
