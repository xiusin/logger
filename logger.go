package logger

import (
	"fmt"
	"os"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var DisableColor = false

type Options struct {
	Level        Level
	RecordCaller bool
	ShortName    bool
}

func DefaultOptions() *Options {
	return &Options{
		Level:        DebugLevel,
		RecordCaller: true,
		ShortName:    true,
	}
}

type AbstractLogger interface {
	SetLogLevel(level Level)

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type Logger struct {
	*writer
	config *Options
}

func New(options *Options) *Logger {
	if options == nil {
		options = DefaultOptions()
	}
	l := &Logger{
		writer: newWriter(os.Stdout, options.RecordCaller),
		config: options,
	}
	return l
}

func (l *Logger) Debug(args ...interface{}) {
	if l.config.Level <= DebugLevel {
		l.writer.WriteString(DebugLevel, fmt.Sprint(args))
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.config.Level <= DebugLevel {
		l.writer.WriteString(DebugLevel, fmt.Sprintf(format, args))
	}
}

func (l *Logger) Print(args ...interface{}) {
	if l.config.Level <= InfoLevel {
		l.writer.WriteString(InfoLevel, fmt.Sprint(args))
	}
}

func (l *Logger) Printf(format string, args ...interface{}) {
	if l.config.Level <= InfoLevel {
		l.writer.WriteString(InfoLevel, fmt.Sprintf(format, args))
	}
}

func (l *Logger) Warning(args ...interface{}) {
	if l.config.Level <= WarnLevel {
		l.writer.WriteString(WarnLevel, fmt.Sprint(args))
	}
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.config.Level <= WarnLevel {
		l.writer.WriteString(WarnLevel, fmt.Sprintf(format, args))
	}
}

func (l *Logger) Error(args ...interface{}) {
	l.writer.WriteString(ErrorLevel, fmt.Sprint(args))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.writer.WriteString(ErrorLevel, fmt.Sprintf(format, args))
}
