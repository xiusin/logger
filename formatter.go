package logger

import "github.com/fatih/color"

type levelFormatter struct {
	ColorType    string
	Type         string
	//LoggerFormat string
}

//const DefaultLoggerFormat = "%date% %type% %file% %message%"

var defaultFormatters = map[Level]*levelFormatter{
	DebugLevel: {
		Type:      "[DBUG]",
		ColorType: color.HiBlueString("%s", "[DBUG]"),
		//LoggerFormat: DefaultLoggerFormat,
	},
	InfoLevel: {
		Type:      "[INFO]",
		ColorType: color.GreenString("[INFO]"),
		//LoggerFormat: DefaultLoggerFormat,
	},
	WarnLevel: {
		Type:      "[WARN]",
		ColorType: color.HiYellowString("[WARN]"),
		//LoggerFormat: DefaultLoggerFormat,
	},
	ErrorLevel: {
		Type:      "[ERRO]",
		ColorType: color.HiRedString("[ERRO]"),
	},
}

//func SetLevelFormatter(level Level, formatter *LevelFormatter) {
//	defaultFormatters[level] = formatter
//}
