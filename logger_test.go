package logger

import (
	"github.com/kataras/golog"
	"io/ioutil"
	"testing"
)

//func TestLogger(t *testing.T) {
//
//	//SetDefault(New(&Options{
//	//	Writer:     os.Stdout,
//	//	Level:      InfoLevel,
//	//	DateFormat: DefaultDateFormat,
//	//}))
//
//	Error("This is an error message")
//	Warning("This is a warning message")
//	Print("This is an info message")
//	Debug("This is a debug message")
//
//	golog.Errorf("This is an error message")
//	golog.Warnf("This is a warning message")
//	golog.Infof("This is an info message")
//	golog.Debugf("This is a debug message")
//}

func BenchmarkLogger(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	f,_ := ioutil.TempFile("", "*")
	b.Log(f.Name())
	SetOutput(f)
	SetLogLevel(DebugLevel)
	for i := 0; i < b.N; i++ {
		Errorf("[%d] This is an error message", i)
		Warningf("[%d] This is a warning message", i)
		Printf("[%d] This is an info message", i)
		Debugf("[%d] This is a debug message", i)
	}
}

func BenchmarkGolog(b *testing.B) {
	// logger defaults
	f,_ := ioutil.TempFile("", "*")
	b.Log(f.Name())
	golog.SetOutput(f)
	golog.SetLevel("debug")
	// disable time formatting because logrus and std doesn't print the time.
	// note that the time is being set-ed to time.Now() inside the golog's Log structure, same for logrus,
	// Therefore we set the time format to empty on golog test in order
	// to acomblish a fair comparison between golog and logrus.
	//golog.SetTimeFormat("")

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		golog.Errorf("[%d] This is an error message", i)
		golog.Warnf("[%d] This is a warning message", i)
		golog.Infof("[%d] This is an info message", i)
		golog.Debugf("[%d] This is a debug message", i)
	}
}
