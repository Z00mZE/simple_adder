package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/adder/conf"
)

var WireSet = wire.NewSet(NewLogger)

const (
	LoggModeProd = `prod`
	LoggModeTest = `test`
	LoggModeDev  = `dev`
)

var presets = map[string]slog.Handler{
	LoggModeProd: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
	LoggModeTest: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}),
	LoggModeDev: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
	),
}

func NewLogger(cfg *conf.Settings) *slog.Logger {
	loggerHandler, isExist := presets[strings.ToLower(cfg.LogLevel)]

	if !isExist {
		loggerHandler = presets[LoggModeTest]
	}

	return slog.New(loggerHandler)
}
