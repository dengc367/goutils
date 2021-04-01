package utils

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger return *zap.SugaredLogger
var Logger *zap.SugaredLogger

// InitLogger return nil
/*
init the sugar Log
*/
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "../test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// XLogger return text
/*
test xlogger
*/
func XLogger(text string) {
	InitLogger()
	defer Logger.Sync() // flushes buffer, if any
	Logger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", text,
		"attempt", 3,
		"backoff", time.Second,
	)
	Logger.Infof("Failed to fetch URL: %s", text)
}
