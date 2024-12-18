package dip_logger

import (
	"dip/bootstrap/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	accessLogger Logger
	appLogger    Logger
	logLock      sync.RWMutex
)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

// Logger is the interface for Logger types
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}

func InitAppLogger() {
	logLock.Lock()
	defer logLock.Unlock()
	accessCore := getZapLogCore(conf.Config.LogLevel, "dip_http_access.log")
	appCore := getZapLogCore(conf.Config.LogLevel, "dip_application.log")
	accessLogger = zap.New(accessCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	appLogger = zap.New(appCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func getZapLogCore(logLevel string, logFile string) zapcore.Core {
	level := getLogLevel(logLevel)
	encoder := getEncoder()
	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), getLogMultiWriter(logFile), level)
}

func getLogLevel(level string) zapcore.Level {
	if zapLevel, ok := levelMap[level]; ok {
		return zapLevel
	}
	return zapcore.InfoLevel
}

func getLogMultiWriter(logFile string) zapcore.WriteSyncer {
	// 确保目录存在，如果不存在则创建
	err := os.MkdirAll(conf.Config.LoggerPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(conf.Config.LoggerPath + logFile)
	if err != nil {
		panic(err)
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stdout))
}

func getEncoder() zapcore.EncoderConfig {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	return encoder
}

// Infof is format info level
func InfofAccess(fmt string, args ...interface{}) {
	if accessLogger != nil {
		accessLogger.Infof(fmt, args...)
	}
}

// Info is info level
func Info(args ...interface{}) {
	if appLogger != nil {
		appLogger.Info(args...)
	}
}

// Warn is warning level
func Warn(args ...interface{}) {
	if appLogger != nil {
		appLogger.Warn(args...)
	}

}

// Error is error level
func Error(args ...interface{}) {
	appLogger.Error(args...)
}

// Debug is debug level
func Debug(args ...interface{}) {
	appLogger.Debug(args...)
}

// Infof is format info level
func Infof(fmt string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Infof(fmt, args...)
	}
}

// Warnf is format warning level
func Warnf(fmt string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Warnf(fmt, args...)
	}

}

// Errorf is format error level
func Errorf(fmt string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Errorf(fmt, args...)
	}

}

// Debugf is format debug level
func Debugf(fmt string, args ...interface{}) {
	if appLogger != nil {
		appLogger.Debugf(fmt, args...)
	}

}
