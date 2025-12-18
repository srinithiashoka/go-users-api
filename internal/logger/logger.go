package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// Init initializes the global zap logger
func Init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	Log = logger
}

// Sync flushes buffered logs (call on app shutdown)
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Helper methods
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}
