package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func New() *zap.Logger {
	config := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		logger = zap.NewNop()
	}
	defer logger.Sync()

	return logger
}

func init() {
	logger = New()
	logger.Sugar().Info("Zap logger initialized")
}

func GetLogger() *zap.Logger {
	return logger
}
