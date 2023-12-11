package elogs

import (
	"log"
	"log/slog"
	"os"
)

type Params struct {
	ServiceName string
	PathToWrite string
	TerminalMsg bool
	LogLevel    int
	RotateSize  int64
}

var term = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func (p *Params) New() *Params {
	return p
}

func (p *Params) fileSize() int64 {
	f, err := os.Stat(p.PathToWrite)
	if err != nil {
		log.Println("Stat ")
	}
	return f.Size()
}

func (p *Params) removeFile() {
	if p.fileSize() >= p.RotateSize {
		err := os.Remove(p.PathToWrite)
		if err != nil {
			log.Println("remove log file", err)
		}
	}

}

func (p *Params) logToFile() *slog.Logger {
	p.removeFile()
	f, err := os.OpenFile(p.PathToWrite, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		term.Error("error opening file", p.PathToWrite, err)
	}
	return slog.New(slog.NewJSONHandler(f, nil))
}

func (p *Params) Info(msg string, args ...any) {
	if p.LogLevel >= 0 {
		args = append(args, "service_name", p.ServiceName)
		if p.TerminalMsg {
			term.Info(msg, args...)
		}
		if p.PathToWrite != "" {
			p.logToFile().Info(msg, args...)
		}
	}
}

func (p *Params) Error(msg string, args ...any) {
	if p.LogLevel >= 1 {
		args = append(args, "service_name", p.ServiceName)
		if p.TerminalMsg {
			term.Error(msg, args...)
		}
		if p.PathToWrite != "" {
			p.logToFile().Error(msg, args...)
		}
	}
}

func (p *Params) Warn(msg string, args ...any) {
	if p.LogLevel >= 2 {
		args = append(args, "service_name", p.ServiceName)
		if p.TerminalMsg {
			term.Warn(msg, args...)
		}
		if p.PathToWrite != "" {
			p.logToFile().Warn(msg, args...)
		}
	}
}
