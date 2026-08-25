package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func SystemLogPath(appName string) string {
	var dir string
	switch runtime.GOOS {
	case "windows":
		programData := os.Getenv("ProgramData")
		if programData == "" {
			programData = `C:\ProgramData`
		}
		dir = filepath.Join(programData, appName, "logs")
	case "darwin":
		dir = filepath.Join("/Library/Logs", appName)
	default:
		dir = filepath.Join("/var/log", appName)
	}

	return filepath.Join(dir, appName+".log")
}

func Setup(logPath string) (*os.File, error) {
	if logPath == "" {
		logPath = SystemLogPath("memap")
	}

	if dir := filepath.Dir(logPath); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return file, nil
}
