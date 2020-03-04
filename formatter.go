package logger

import "github.com/fatih/color"

type LevelFormatter struct {
	ColorType       string
	Type            string
	DateFormat      string
	LoggerFormatter string
}


var defaultFormatters = map[Level]*LevelFormatter{
	DebugLevel: {
		Type:            "[DEBUG]",
		ColorType:       color.HiBlueString("[DEBUG]"),
		DateFormat:      "2006-01-02 15:04:05",
		LoggerFormatter: "%date% %type% %file% %message%",
	},
	InfoLevel: {
		Type:            "[INFO]",
		ColorType:       color.GreenString("[INFO]"),
		DateFormat:      "2006-01-02 15:04:05",
		LoggerFormatter: "%date% %type %file% %message%",
	},
	WarnLevel: {
		Type:            "[WARNING]",
		ColorType:       color.HiBlueString("[WARNING]"),
		DateFormat:      "2006-01-02 15:04:05",
		LoggerFormatter: "%date% %type %file% %message%",
	},
	ErrorLevel: {
		Type:            "[ERROR]",
		ColorType:       color.HiBlueString("[ERROR]"),
		DateFormat:      "2006-01-02 15:04:05",
		LoggerFormatter: "%date% %type% %file% %message%",
	},
}

func SetLevelFormatter(level Level, formatter *LevelFormatter) {
	defaultFormatters[level] = formatter
}