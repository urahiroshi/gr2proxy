package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LocalPort int
	AdminPort int
	Logger    *zap.Logger
}

func NewLogger(level zapcore.Level) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	return config.Build()
}
