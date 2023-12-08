package elogs

import (
	"log/slog"
	"os"
)

type LoggerParams struct {
	ServiceName string
	PathToWrite string
	TerminalMsg bool
	LogLevel    int
}

var term = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func logToFile(path string) *slog.Logger {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		term.Error("error opening file", path, err)
	}
	return slog.New(slog.NewJSONHandler(f, nil))
}

func (l *LoggerParams) Info(msg string, args ...any) {
	if l.LogLevel == 0 || l.LogLevel == 1 || l.LogLevel <= 2 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Info(msg, args...)
		}
		if l.PathToWrite != "" {
			logToFile(l.PathToWrite).Info(msg, args...)
		}
	}
}

func (l *LoggerParams) Error(msg string, args ...any) {
	if l.LogLevel == 0 || l.LogLevel == 1 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Error(msg, args...)
		}
		if l.PathToWrite != "" {
			logToFile(l.PathToWrite).Error(msg, args...)
		}
	}
}

func (l *LoggerParams) Warn(msg string, args ...any) {
	if l.LogLevel <= 2 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Warn(msg, args...)
		}
		if l.PathToWrite != "" {
			logToFile(l.PathToWrite).Warn(msg, args...)
		}
	}
}
