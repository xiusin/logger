// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package zap

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"github.com/xiusin/pine/logger"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
	config *Options
}

type Options struct {
	TimeFormat         string
	Level              logger.Level
	RotateLogDirFormat string
	LogDir             string
	DebugLogName       string
	InfoLogName        string
	WarnLogName        string
	ErrorLogName       string
	Console            bool
	MaxSizeMB          int
	MaxBackups         int
	MaxAgeDay          int
	Compress           bool // 压缩日志.(分割时)
}

func DefaultOptions() *Options {
	return &Options{
		TimeFormat:         "2006-01-02 15:04:05",
		Level:              logger.DebugLevel,
		RotateLogDirFormat: "2006-01-02",
		DebugLogName:       "debug.log",
		InfoLogName:        "info.log",
		WarnLogName:        "warn.log",
		ErrorLogName:       "error.log",
		Console:            true,
		MaxAgeDay:          7,
		MaxSizeMB:          50, //50M
		MaxBackups:         3,
		Compress:           true,
	}
}

func New(options *Options) *Logger {
	if options == nil {
		options = DefaultOptions()
	}
	if viper.GetInt("env") == 0 {
		options.Console = true
	}
	infoLevelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return zapcore.InfoLevel >= zapcore.Level(options.Level)
	})
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { //时间格式编码器
			enc.AppendString(t.Format(options.TimeFormat))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	errLevelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return zapcore.InfoLevel < zapcore.Level(options.Level)
	})
	var core zapcore.Core
	if !options.Console {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(writer(options.InfoLogName, options)), infoLevelEnabler),
			zapcore.NewCore(encoder, zapcore.AddSync(writer(options.ErrorLogName, options)), errLevelEnabler),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(writer(options.InfoLogName, options)), infoLevelEnabler),
			zapcore.NewCore(encoder, zapcore.AddSync(writer(options.ErrorLogName, options)), errLevelEnabler),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level > zapcore.Level(options.Level)
			})),
		)
	}

	return &Logger{Logger: zap.New(core, zap.AddCaller()), config: options}
}

func (l *Logger) Debug(msg string, args ...interface{}) {


	//l.Logger.Debug(msg, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Print(msg string, args ...interface{}) {
	l.Logger.Info(msg)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warning(msg string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Logger.Error(msg)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(msg, args...))
}

func writer(filename string, option *Options) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath.Join(option.LogDir, option.RotateLogDirFormat, filename),
		MaxSize:    option.MaxSizeMB,
		MaxBackups: option.MaxBackups,
		MaxAge:     option.MaxAgeDay,
		Compress:   option.Compress,
	}
}
