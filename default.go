package logger

import "io"

var defaultLogger AbstractLogger = New()

func init()  {
	defaultLogger.SetReportCaller(true, 4)
}

func SetDefault(logger AbstractLogger) {
	defaultLogger = logger
}

func SetLogLevel(level Level) {
	defaultLogger.SetLogLevel(level)
}

func SetOutput(writer io.Writer) {
	defaultLogger.SetOutput(writer)
}

func SetDateFormat(format string) {
	defaultLogger.SetDateFormat(format)
}

func SetReportCaller(b bool, skipCallerNumber ...int) {
	defaultLogger.SetReportCaller(b, skipCallerNumber...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}
