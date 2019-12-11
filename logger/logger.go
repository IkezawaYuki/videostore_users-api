package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init(){
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		EncoderConfig:     zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller:zapcore.ShortCallerEncoder,
		},
	}
	var err error
	log, err = logConfig.Build()
	if err != nil{
		panic(err)
	}
}

func GetLogger()*zap.Logger{
	return log
}

func Info(msg string, tags ...zap.Field){
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field){
	if err != nil{
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	log.Sync()
}