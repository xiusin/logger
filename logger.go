package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

const DefaultDateFormat = "2006/01/02 15:04"

const DefaultSkipCallerNumber = 3

var (
	risk  = []byte(":")
	left  = []byte("[")
	right = []byte("]")
	space = []byte(" ")
	brk = []byte("\n")
	DisableColor = false

	bufPool sync.Pool
)

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

func init() {
	bufPool = sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(nil)
	}}
}

func New() *Logger {
	return &Logger{
		Writer:           os.Stdout,
		Level:            DebugLevel,
		DateFormat:       DefaultDateFormat,
		SkipCallerNumber: DefaultSkipCallerNumber,
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
		l.SkipCallerNumber = DefaultSkipCallerNumber
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
	bytesBuf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		bytesBuf.Reset()
		bufPool.Put(bytesBuf)
	}()
	if len(l.DateFormat) > 0 {
		bytesBuf.Write(left)
		bytesBuf.WriteString(time.Now().Format(l.DateFormat))
		bytesBuf.Write(right)
		bytesBuf.Write(space)
	}
	bytesBuf.WriteString(t)
	bytesBuf.Write(space)
	l.writeCallerInfo(bytesBuf)
	bytesBuf.WriteString(message)
	bytesBuf.Write(brk)
	l.Writer.Write(bytesBuf.Bytes())
}

func (l *Logger) writeCallerInfo(buf *bytes.Buffer) {
	if l.RecordCaller {
		_, callerFile, line, ok := runtime.Caller(l.SkipCallerNumber)
		if ok {
			buf.Write([]byte(path.Base(callerFile)))
			buf.Write(risk)
			buf.WriteString(strconv.Itoa(line))
			buf.Write(risk)
			buf.Write(space)
		}
	}
}
