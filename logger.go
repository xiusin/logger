package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var DisableColor = false
const DefaultDateFormat = "2006/01/02 15:04"

type AbstractLogger interface {
	SetLogLevel(level Level)
	SetOutput(writer io.Writer)
	SetReportCaller(b bool, skipCallerNumber ...int)
	SetDateFormat(format string)

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
	io.Writer
	Level            Level
	DateFormat       string
	RecordCaller     bool
	SkipCallerNumber int
}

func New() *Logger {
	return &Logger{
		Writer:           os.Stdout,
		Level:            InfoLevel,
		DateFormat:       DefaultDateFormat,
		SkipCallerNumber: 3,
	}
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.Writer = writer
}

func (l *Logger) SetLogLevel(level Level) {
	l.Level = level
}

func (l *Logger) SetDateFormat(format string) {
	l.DateFormat = format
}

func (l *Logger) SetReportCaller(b bool, skipCallerNumber ...int) {
	l.RecordCaller = b
	if len(skipCallerNumber) == 0 {
		l.SkipCallerNumber = 3
	} else if skipCallerNumber[0] > 0 {
		l.SkipCallerNumber = skipCallerNumber[0]
	}
}

func (l *Logger) Debug(args ...interface{}) {
	if l.Level <= DebugLevel {
		l.WriteString(DebugLevel, fmt.Sprint(args...))
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Level <= DebugLevel {
		l.WriteString(DebugLevel, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Print(args ...interface{}) {
	if l.Level <= InfoLevel {
		l.WriteString(InfoLevel, fmt.Sprint(args...))
	}
}

func (l *Logger) Printf(format string, args ...interface{}) {
	if l.Level <= InfoLevel {
		l.WriteString(InfoLevel, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warning(args ...interface{}) {
	if l.Level <= WarnLevel {
		l.WriteString(WarnLevel, fmt.Sprint(args...))
	}
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.Level <= WarnLevel {
		l.WriteString(WarnLevel, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(args ...interface{}) {
	l.WriteString(ErrorLevel, fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.WriteString(ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *Logger) WriteString(level Level, message string) {
	t := defaultFormatters[level].ColorType
	if DisableColor {
		t = defaultFormatters[level].Type
	}
	var bytesBuf = bytes.NewBuffer(nil)
	if len(l.DateFormat) > 0 {
		bytesBuf.WriteString("[")
		bytesBuf.WriteString(time.Now().Format(l.DateFormat))
		bytesBuf.WriteString("] ")
	}
	bytesBuf.WriteString(l.getCaller())
	bytesBuf.WriteString(t)
	bytesBuf.WriteString(" ")
	bytesBuf.WriteString(message)
	bytesBuf.WriteString("\n")

	l.Writer.Write(bytesBuf.Bytes())
}

func (l *Logger) getCaller() string {
	if l.RecordCaller {
		_, callerFile, line, ok := runtime.Caller(l.SkipCallerNumber)
		if ok {
			return path.Base(callerFile) + ":" + strconv.Itoa(line) + ": "
		}
	}
	return ""
}
