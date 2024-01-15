package logger

import (
	"log/slog"
	"os"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(NewLogger)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
