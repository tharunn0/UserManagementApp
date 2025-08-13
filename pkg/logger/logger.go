package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logr *zap.Logger

func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	logger, er := cfg.Build()
	if er != nil {
		log.Fatalln(er)
	}
	logr = logger
}

func Info(msg string, fields ...zap.Field) {
	logr.Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	logr.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	logr.Error(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	logr.Panic(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	logr.Fatal(msg, fields...)
}
