package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
我们先定义一个包级的全局变量l，类型是*zap.Logger，并创建了InitLogger和GetLogger这两个函数。
因为，zap不支持日志归档，所以在InitLogger中定义了一个lumberjack的hook，用来归档日志。
*/

var l *zap.Logger

func InitLogger(logPath, logLevel string) error {
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1024,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
	w := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level)

	l = zap.New(core)
	return nil
}

func GetLogger() *zap.Logger {
	return l
}
