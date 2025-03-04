package logger

import (
	"log/slog"
	"os"
)

var LogLevels = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

var Logger *slog.LevelVar

func SetUpLogger(logLevel string) {
	programLevel := new(slog.LevelVar)

	level, exists := LogLevels[logLevel]
	if !exists {
		level = slog.LevelWarn
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	programLevel.Set(level)
	slog.SetDefault(logger)
	Logger = programLevel
}
