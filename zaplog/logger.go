package zaplog

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var onlyOnce sync.Once

var Logger *zap.SugaredLogger

func InitializeLogger() {
	onlyOnce.Do(func() {
		initLogger()
	})
}

func initLogger() {
	logpath := "lite.log"
	configLogger := zap.NewProductionEncoderConfig()
	// configLogger.EncodeLevel = zapcore.CapitalColorLevelEncoder
	configLogger.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	fileEncoderLogger := zapcore.NewConsoleEncoder(configLogger)
	consoleEncoder := zapcore.NewConsoleEncoder(configLogger)
	// writer := zapcore.AddSync(wLoggerRotate)
	fileLogger, _ := os.OpenFile(logpath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(fileLogger)
	defaultLogLevel := zapcore.InfoLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoderLogger, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
}
