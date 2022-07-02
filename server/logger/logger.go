package logger

import (
	"fmt"
	"os"

	"github.com/kataras/golog"
)

func init() {
	golog.Levels[golog.ErrorLevel].Title = "[ERROR]"
	golog.Levels[golog.DebugLevel].Title = "[DEBUG]"
}

type Logger interface {
	Error(v ...any)
	Errorf(format string, args ...any)

	Warn(v ...any)
	Warnf(format string, args ...any)

	Info(v ...any)
	Infof(format string, args ...any)

	Debug(v ...any)
	Debugf(format string, args ...any)
}

type Log struct {
	outputFile *os.File

	*golog.Logger
}

const logTimeFormat = "02.01.2006 15:04:05 MST"

func NewLog(log *golog.Logger) *Log {
	log.TimeFormat = logTimeFormat

	return &Log{
		Logger: log,
	}
}

func (l *Log) SetOutputFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}

	l.outputFile = file
	l.SetOutput(l.outputFile)

	return nil
}

func (l *Log) Close() error {
	if l.outputFile != nil {
		return l.outputFile.Close()
	}

	return nil
}
