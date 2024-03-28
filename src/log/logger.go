package log

import (
	"log/slog"
	"time"

	"github.com/KenshiTech/unchained/config"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

var Logger *slog.Logger
var Levels map[string]slog.Level

func init() {
	Levels = make(map[string]slog.Level)
	Levels["info"] = slog.LevelInfo
	Levels["warn"] = slog.LevelWarn
	Levels["debug"] = slog.LevelDebug
	Levels["error"] = slog.LevelError
}

func Start() {
	Logger = slog.New(tint.NewHandler(colorable.NewColorableStdout(), &tint.Options{
		Level:      Levels[config.App.System.Log],
		TimeFormat: time.RFC3339,
	}))
}
