package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func zap1() {
	fmt.Println("========== zap1() Preset Development Logger ==========")
	logger := newPresetDevelopmentLogger()
	defer logger.Sync()
	tryLog(logger)
}

func zap2() {
	fmt.Println("========== zap2) Preset Production Logger ==========")
	logger := newPresetProductionLogger()
	defer logger.Sync()
	tryLog(logger)
}

func zap3() {
	fmt.Println("========== zap3() Custom Logger ==========")
	logger := newCustomLogger()
	defer logger.Sync()

	for i := 0; i < 1000; i++ {
		tryLog(logger)
	}
}

func newPresetDevelopmentLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func newPresetProductionLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func newCustomLogger() *zap.Logger {
	writeSyncer := getRotatedLogWriter()
	encoder := getJSONEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getCustomJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

func getCustomConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func getRotatedLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func tryLog(logger *zap.Logger) {
	const url = "http://example.com"

	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
