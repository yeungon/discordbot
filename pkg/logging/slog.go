package logging

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var logFile *os.File

// Log sets up the global logger with a JSON handler.
func Log() {
	var err error
	const logPath = "/app/logs/discordbot.json"
	// Ensure the logs/ directory exists
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create a JSON handler that writes to the file
	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Capture debug messages and above
	}))

	// Set the logger as the default so that slog uses it
	slog.SetDefault(logger)
}

// CloseLog ensures that the log file is properly closed and flushed.
func CloseLog() {
	if logFile != nil {
		// Flush any buffered data to the file
		if err := logFile.Sync(); err != nil {
			log.Printf("Failed to flush log file: %v\n", err)
		}
		// Close the log file
		if err := logFile.Close(); err != nil {
			log.Printf("Failed to close log file: %v\n", err)
		}
	}
}
