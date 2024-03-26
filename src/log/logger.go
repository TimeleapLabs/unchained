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

func Start() error {
	// TODO fix the log package to not be singleton.
	Logger = slog.New(tint.NewHandler(colorable.NewColorableStdout(), &tint.Options{
		Level:      Levels[config.Config.GetString("log")],
		TimeFormat: time.RFC3339,
	}))
	return nil
}
