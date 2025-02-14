package utils

import (
	"log"
	"os"
	"time"

	"app-api/pkg/zaplogger"

	"go.uber.org/zap"
)

type zapLog struct{}

func ZapLog() *zapLog {
	return &zapLog{}
}

func (s *zapLog) Info(filename string, message string, args ...zap.Field) {
	s.Write(filename, "info", message, args...)
}

func (s *zapLog) Debug(filename string, message string, args ...zap.Field) {
	s.Write(filename, "debug", message, args...)
}

func (s *zapLog) Warn(filename string, message string, args ...zap.Field) {
	s.Write(filename, "warn", message, args...)
}

func (s *zapLog) Error(filename string, message string, args ...zap.Field) {
	s.Write(filename, "error", message, args...)
}

func (s *zapLog) Write(filename string, level string, message string, args ...zap.Field) {
	directoryName := "./storege/logs/" + filename
	err := os.MkdirAll(directoryName, 0755) // 创建目录并设置权限
	if err != nil {
		log.Fatal("创建日志目录错误:", err)
		return
	}

	logFilename := directoryName + "/" + filename + "-" + time.Now().Format("2006-01-02") + ".log"
	// file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("创建文件错误:", err)
	// }

	var zapLevel zaplogger.Option
	switch level {
	case "info":
		zapLevel = zaplogger.WithInfoLevel()
	case "debug":
		zapLevel = zaplogger.WithDebugLevel()
	case "warn":
		zapLevel = zaplogger.WithWarnLevel()
	case "error":
		zapLevel = zaplogger.WithErrorLevel()
	default:
		zapLevel = zaplogger.WithInfoLevel()
	}

	logger, err := zaplogger.NewJSONLogger(
		zapLevel,
		zaplogger.WithFileRotationP(logFilename),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	switch level {
	case "info":
		logger.Info(message, args...)
	case "debug":
		logger.Debug(message, args...)
	case "warn":
		logger.Warn(message, args...)
	case "error":
		logger.Error(message, args...)
	default:
		logger.Info(message, args...)
	}

}
