package logger

import (
	"io"
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	// 1. Create a log file
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// 2. Define multiple outputs
	// Console: Pretty text | File: JSON for production analysis
	multiWriter := io.MultiWriter(os.Stdout, file)

	// 3. Configure Handler (JSON for easy parsing by tools like ELK or Datadog)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // Log everything from Debug up to Error
	}

	handler := slog.NewJSONHandler(multiWriter, opts)

	// 4. Create the logger
	logger := slog.New(handler)

	// Set as global logger (optional)
	slog.SetDefault(logger)

	return logger
}
