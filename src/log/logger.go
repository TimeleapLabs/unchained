package log

import (
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

var Logger *slog.Logger

func init() {
	Logger = slog.New(tint.NewHandler(colorable.NewColorableStdout(), &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
	}))
}
