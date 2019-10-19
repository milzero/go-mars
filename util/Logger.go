package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var Logger *zap.SugaredLogger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", "./switcher.log"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	Logger = logger.Sugar()
}