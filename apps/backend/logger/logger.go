package logger

import (
	"go.uber.org/zap"
	"medivu.co/auth/envs"
)

var logger *zap.Logger

func Init() {
	var err error

	if !envs.IsProduction() {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.CallerKey = "" // Caller 정보 비활성화
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig = encoderConfig
		logger, err = config.Build()
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.CallerKey = ""
		config := zap.NewProductionConfig()
		config.EncoderConfig = encoderConfig
		logger, err = config.Build()
		
	}
	if err != nil {
		panic(err)
	}
}

// this function may panic if Init() was not called before
func Get() *zap.Logger {
	if logger == nil {
		panic("Logger not initialized. Call Init() first.")
	}
	return logger
}

func Sync() {
	if err := logger.Sync(); err != nil {
		panic(err)
	}
}
