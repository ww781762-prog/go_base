package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type ConsoleLogger struct {
	LogLevel    int
	File        *os.File
	WFile       *os.File
	LogDataChan chan *LogData
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
	//fmt.Fprintf(file, "%s [%s] %s %s:%d %s\n", nowStr, leveltest, filename, funname, lineon, msg)

	logstr := fmt.Sprintf("%s [%s] %s %s:%d %s\n", nowStr, leveltest, filename, funname, lineon, msg)
	ld := &LogData{
		file:   file,
		string: logstr,
	}
	select {
	case cl.LogDataChan <- ld:
	default:
	}
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
func (cl *ConsoleLogger) writeLogBackGroud() {
	for ld := range cl.LogDataChan {
		// 模拟写入卡顿
		time.Sleep(10 * time.Second)
		_, err := fmt.Fprintf(ld.file, "%s", ld.string)
		if err != nil {
			fmt.Printf("writeLogBackGroud 写日志失败%v\n", err)
		}
	}
}

func (cl *ConsoleLogger) Init() {
	cl.File = os.Stdout
	cl.WFile = os.Stderr
	go cl.writeLogBackGroud()
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

func NewConsoleLogger(config map[string]string) (log LoggerInterface, err error) {

	logLevel, ok := config["log_level"]
	if !ok {
		logLevel = "info"
	}
	logint := GetLevelInt(logLevel)
	log = &ConsoleLogger{
		LogLevel:    logint,
		LogDataChan: make(chan *LogData, 5000),
	}
	log.Init()
	return log, nil
}
