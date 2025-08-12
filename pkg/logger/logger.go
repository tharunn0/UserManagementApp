package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logr *zap.Logger

func InitLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, er := cfg.Build()
	if er != nil {
		log.Fatalln(er)
	}
	defer logger.Sync()
	Logr = logger
	return logger
}

func ZapLogger(lgr *zap.Logger) {

}
