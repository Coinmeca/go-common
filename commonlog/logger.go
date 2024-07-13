package commonlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	UseTerminal       bool
	UseFile           bool
	VerbosityTerminal int
	VerbosityFile     int
	FilePath          string
}

var Logger *zap.Logger

func InitLog(filePath string, useConsole bool) {
	var cores []zapcore.Core

	if useConsole {
		consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
		cores = append(cores, consoleCore)
	}

	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel)
	cores = append(cores, fileCore)

	combinedCore := zapcore.NewTee(cores...)
	Logger = zap.New(combinedCore)
	zap.ReplaceGlobals(Logger)
}
