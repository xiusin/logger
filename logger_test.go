package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Error("This is an error message")
	Warning("This is a warning message")
	Print("This is an info message")
	Debug("This is a debug message")
}

func BenchmarkLogger(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	SetLogLevel(DebugLevel)
	for i := 0; i < b.N; i++ {
		Errorf("[%d] This is an error message", i)
		Warningf("[%d] This is a warning message", i)
		Printf("[%d] This is an info message", i)
		Debugf("[%d] This is a debug message", i)
	}
}

