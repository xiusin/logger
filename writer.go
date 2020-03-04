package logger

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"
	"time"
)

type writer struct {
	writers      io.Writer
	recordCaller bool
}

func newWriter(w io.Writer, recordCaller bool) *writer {
	return &writer{w, recordCaller}
}

func (w *writer) WriteString(level Level, message string) {
	formatter := defaultFormatters[level]
	storeMessage := formatter.LoggerFormatter
	t := formatter.ColorType
	if DisableColor {
		t = formatter.Type
	}

	storeMessage = strings.Replace(storeMessage, "%date%", time.Now().Format(formatter.DateFormat), 1)
	storeMessage = strings.Replace(storeMessage, "%type%", t, 1)
	storeMessage = strings.Replace(storeMessage, "%file%", w.getCaller(), 1)
	storeMessage = strings.Replace(storeMessage, "%message%", message, 1)

	w.writers.Write([]byte(storeMessage))
}

func (l *writer) getCaller() string {
	if l.recordCaller {
		_, callerFile, line, ok := runtime.Caller(3)
		if ok {
			return fmt.Sprintf("%s:%d:", path.Base(callerFile), line)
		}
	}
	return ""
}
