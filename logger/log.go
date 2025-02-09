package logger

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TODO: other domains logger update
type Config struct {
	UseTerminal       bool
	UseFile           bool
	VerbosityTerminal int
	VerbosityFile     int
	FilePath          string
}

func InitLog(config Config) {
    var handlers []log.Handler

    if config.UseTerminal {
        handler := log.LvlFilterHandler(config.VerbosityTerminal,
            log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
        handlers = append(handlers, handler)
    }

    if config.UseFile {
        handler := log.LvlFilterHandler(config.VerbosityFile, log.StreamHandler(&lumberjack.Logger{
            Filename:   config.FilePath,
            MaxSize:    64,
            MaxBackups: 3,
            MaxAge:     28,   // days
            Compress:   true, // disabled by default
        }, log.JSONFormat()))
        handlers = append(handlers, handler)
    }

    log.Root().SetHandler(log.MultiHandler(handlers...))
}

// Trace is a convenient alias for Root().Trace
func Trace(msg string, ctx ...interface{}) {
	log.Trace(msg, ctx...)
}

// Debug is a convenient alias for Root().Debug
func Debug(msg string, ctx ...interface{}) {
	log.Debug(msg, ctx...)
}

// Info is a convenient alias for Root().Info
func Info(msg string, ctx ...interface{}) {
	log.Info(msg, ctx...)
}

// Warn is a convenient alias for Root().Warn
func Warn(msg string, ctx ...interface{}) {
	log.Warn(msg, ctx...)
}

// Error is a convenient alias for Root().Error
func Error(msg string, ctx ...interface{}) {
	log.Error(msg, ctx...)
}

// Crit is a convenient alias for Root().Crit
func Crit(msg string, ctx ...interface{}) {
	// utils.SendChatAlert(msg)
	log.Crit(msg, ctx...)
}
