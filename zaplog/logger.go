package zaplog

import (
	"context"
	"firstwails/domain"
	"os"
	"path"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var onlyOnce sync.Once

var LoggerShugar *zap.SugaredLogger
var Reductor *zap.Logger
var EchoSugar *zap.Logger

func OnStartup() {
	onlyOnce.Do(func() {
		initLogger()
	})
}

func OnShutdown() {
	LoggerShugar.Sync()
	Reductor.Sync()
	EchoSugar.Sync()
}

// возврат только после прерывания контекста
func Run(ctx context.Context) error {
	ch := make(chan struct{})
	go func() {
		<-ctx.Done()
		ch <- struct{}{}
	}()
	<-ch
	return nil
}

func run(ctx context.Context) error {
	// если завершатется контекст из майн то и эти все задачи завершаются
	go func() {
		<-ctx.Done()
		LoggerShugar.Info("zaplog OnShutdown!")
		OnShutdown()
	}()
	return nil
}

func initLogger() {
	logpath := path.Join(domain.LogPath, "lite.log")
	configLogger := zap.NewProductionEncoderConfig()
	// configLogger.EncodeLevel = zapcore.CapitalColorLevelEncoder
	configLogger.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	fileEncoderLogger := zapcore.NewConsoleEncoder(configLogger)
	consoleEncoder := zapcore.NewConsoleEncoder(configLogger)
	// writer := zapcore.AddSync(wLoggerRotate)
	fileLogger, _ := os.OpenFile(logpath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(fileLogger)
	// defaultLogLevel := zapcore.InfoLevel
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoderLogger, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	LoggerShugar = zap.New(core, zap.AddCaller()).Sugar()
	configReductor := zap.NewProductionEncoderConfig()
	configReductor.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logpathReductor := path.Join(domain.LogPath, "reductor.log")
	reductorLogFile, _ := os.OpenFile(logpathReductor, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fileReductorEncoder := zapcore.NewConsoleEncoder(configReductor)
	writerReductor := zapcore.AddSync(reductorLogFile)
	core2 := zapcore.NewTee(
		zapcore.NewCore(fileReductorEncoder, writerReductor, defaultLogLevel),
	)
	Reductor = zap.New(core2, zap.AddCaller())

	configEcho := zap.NewProductionEncoderConfig()
	configEcho.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logpathEcho := path.Join(domain.LogPath, "echo.log")
	echoLogFile, _ := os.OpenFile(logpathEcho, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fileEchoEncoder := zapcore.NewConsoleEncoder(configEcho)
	writerEcho := zapcore.AddSync(echoLogFile)
	core5 := zapcore.NewTee(
		zapcore.NewCore(fileEchoEncoder, writerEcho, defaultLogLevel),
	)
	EchoSugar = zap.New(core5, zap.AddCaller())
}
