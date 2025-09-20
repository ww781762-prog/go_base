package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type ConsoleLogger struct {
	LogLevel int
	File     *os.File
	WFile    *os.File
}

func (cl *ConsoleLogger) Write(file *os.File, level int, format string, args ...interface{}) {
	if level > cl.LogLevel {
		return
	}
	new := time.Now()
	nowStr := new.Format("2006-01-02 15:04:05")
	filename, funname, lineon := GetLineInfo()
	filename = path.Base(filename)
	funname = path.Base(funname)
	leveltest := GetLevelText(level)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s [%s] %s %s:%d %s\n", nowStr, leveltest, filename, funname, lineon, msg)
}

func (cl *ConsoleLogger) Debug(format string, args ...interface{}) {
	cl.Write(cl.File, LevelDebug, format, args...)

}

func (cl *ConsoleLogger) Info(format string, args ...interface{}) {
	cl.Write(cl.File, LevelInfo, format, args...)
}

func (cl *ConsoleLogger) Warn(format string, args ...interface{}) {
	cl.Write(cl.WFile, LevelWarn, format, args...)
}

func (cl *ConsoleLogger) Error(format string, args ...interface{}) {
	cl.Write(cl.WFile, LevelError, format, args...)
}

func (cl *ConsoleLogger) Fatal(format string, args ...interface{}) {
	cl.Write(cl.WFile, LevelFatal, format, args...)
}

func (cl *ConsoleLogger) Close() {
	cl.File.Close()
	cl.WFile.Close()

}

func (cl *ConsoleLogger) Init() {
	cl.File = os.Stdout
	cl.WFile = os.Stderr

}
func (cl *ConsoleLogger) setLevel(level int) {
	if level < cl.LogLevel {
		cl.LogLevel = LevelFatal
	} else if level > cl.LogLevel {
		cl.LogLevel = LevelDebug
	} else {
		cl.LogLevel = level
	}
}

func NewConsoleLogger(map[string]string) (log LoggerInterface, err error) {
	log = &ConsoleLogger{}
	log.Init()
	return log, nil
}
