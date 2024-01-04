package logger

import (
	"log/slog"
	"os"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(NewSLogger)

func NewSLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
