package utils

import (
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

var Logger *slog.Logger

// Start initializes the logger.
func SetupLogger(logLevel string) {
	levels := make(map[string]slog.Level)
	levels["info"] = slog.LevelInfo
	levels["warn"] = slog.LevelWarn
	levels["debug"] = slog.LevelDebug
	levels["error"] = slog.LevelError

	Logger = slog.New(tint.NewHandler(colorable.NewColorableStdout(), &tint.Options{
		Level:      levels[logLevel],
		TimeFormat: time.RFC3339,
	}))

	Logger.With("Level", logLevel).Info("Logger started")
}
