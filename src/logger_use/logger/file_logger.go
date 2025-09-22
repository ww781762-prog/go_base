package logger

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileLogger struct {
	LogPath           string
	LogLevel          int
	LogName           string
	File              *os.File
	WFile             *os.File
	LogDataChan       chan *LogData
	SplitType         string
	SplitIntervalTime int64
	SplitSize         int64
	SplitLastTime     int64
}

func (fl *FileLogger) Write(file *os.File, level int, format string, args ...interface{}) {
	//fmt.Println(level, fl.LogLevel)

	if level > fl.LogLevel {
		return
	}
	new := time.Now()
	nowStr := new.Format("2006-01-02 15:04:05")
	filename, funname, lineon := GetLineInfo()
	filename = path.Base(filename)
	funname = path.Base(funname)
	leveltest := GetLevelText(level)
	msg := fmt.Sprintf(format, args...)

	// 如果我们直接写日志到文件会出现的问题，如果磁盘有问题会导致整个业务流程有问题，所以我们需要异步去写日志。
	//fmt.Fprintf(file, "%s [%s] %s %s:%d %s\n", nowStr, leveltest, filename, funname, lineon, msg)
	// 并发写
	logStr := fmt.Sprintf("%s [%s] %s %s:%d %s\n", nowStr, leveltest, filename, funname, lineon, msg)
	ld := &LogData{
		file:     file,
		string:   logStr,
		isErrLog: level < 2,
	}

	select {
	case fl.LogDataChan <- ld:
	default:
	}

}

func (fl *FileLogger) writeLogBackGroud() {
	for ld := range fl.LogDataChan {
		if ld.isErrLog {
			if ld.file != fl.WFile {
				ld.file = fl.WFile
			}
		} else {
			if ld.file != fl.File {
				ld.file = fl.File
			}
		}
		fileInfo, err := ld.file.Stat()
		if err != nil {
			fmt.Printf("日志切割，获取日志文件信息失败：%v\n", err)
		}
		fileName := fileInfo.Name()
		if fl.SplitType == "size" {
			if fileInfo.Size() > fl.SplitSize*1024*1024 {
				f1 := fmt.Sprintf("%s/%s", fl.LogPath, fileName)
				backFile := fmt.Sprintf("%s/%s_%s", fl.LogPath, fileName, time.Now().Format("2006-01-02_15:04:05"))
				err = os.Rename(f1, backFile)
				if err != nil {
					fmt.Println("重名名日志文件出错")
				}
				_ = ld.file.Close()

				file, err := os.OpenFile(f1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println("日志切割，重新打开文件失败: ", f1)
				}
				if strings.HasSuffix(fileName, "warn.logs") {
					fl.WFile = file
				} else {
					fl.File = file
				}
				ld.file = file
			}
		}

		if fl.SplitType == "time" {
			if time.Now().Unix()-fl.SplitLastTime > fl.SplitIntervalTime*60 {
				f1 := fmt.Sprintf("%s/%s", fl.LogPath, fileName)
				backFile := fmt.Sprintf("%s/%s_%s", fl.LogPath, fileName, time.Now().Format("2006-01-02_15:04:05"))
				err = os.Rename(f1, backFile)
				_ = ld.file.Close()
				if err != nil {
					fmt.Println("重名名日志文件出错")
				}
				_ = ld.file.Close()

				file, err := os.OpenFile(f1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println("日志切割，重新打开文件失败: ", f1)
				}
				if strings.HasSuffix(fileName, "warn.logs") {
					fl.WFile = file
				} else {
					fl.File = file
				}
				ld.file = file
			}
		}
		_, err = fmt.Fprintf(ld.file, "%s", ld.string)
		if err != nil {
			fmt.Printf("writeLogBackGroud 写日志失败%v\n", err)
		}
	}
}

func (fl *FileLogger) Debug(format string, args ...interface{}) {
	fl.Write(fl.File, LevelDebug, format, args...)

}

func (fl *FileLogger) Info(format string, args ...interface{}) {
	fl.Write(fl.File, LevelInfo, format, args...)
}

func (fl *FileLogger) Warn(format string, args ...interface{}) {
	fl.Write(fl.File, LevelInfo, format, args...)
	fl.Write(fl.WFile, LevelWarn, format, args...)
}

func (fl *FileLogger) Error(format string, args ...interface{}) {
	fl.Write(fl.File, LevelInfo, format, args...)
	fl.Write(fl.WFile, LevelError, format, args...)
}

func (fl *FileLogger) Fatal(format string, args ...interface{}) {
	fl.Write(fl.File, LevelInfo, format, args...)
	fl.Write(fl.WFile, LevelFatal, format, args...)
}

func (fl *FileLogger) Close() {
	fl.File.Close()
	fl.WFile.Close()

}

func (fl *FileLogger) Init() {
	f1 := fmt.Sprintf("%s/%s.logs", fl.LogPath, fl.LogName)
	file, err := os.OpenFile(f1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("fail to open logs file ")
	}
	fw := fmt.Sprintf("%s/%s.warn.logs", fl.LogPath, fl.LogName)
	wFile, err := os.OpenFile(fw, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("fail to open err logs file ")
	}
	fl.File = file
	fl.WFile = wFile
	go fl.writeLogBackGroud()
}
func (fl *FileLogger) setLevel(level int) {
	if level < fl.LogLevel {
		fl.LogLevel = LevelFatal
	} else if level > fl.LogLevel {
		fl.LogLevel = LevelDebug
	} else {
		fl.LogLevel = level
	}
}

//func NewFileLogger(logPath string, logName string, logLeve int) LoggerInterface {
//	file := &FileLogger{
//		LogPath:  logPath,
//		LogLevel: logLeve,
//		LogName:  logName,
//	}
//	file.init()
//	return file
//}

func NewFileLogger(config map[string]string) (log LoggerInterface, err error) {
	LogPath, ok := config["log_path"]
	if !ok {
		err := fmt.Errorf("fail to load log path %s", LogPath)
		return nil, err
	}
	logName, ok := config["log_name"]
	if !ok {
		err := fmt.Errorf("fail to load log name %s", logName)
		return nil, err
	}
	logLevel, ok := config["log_level"]
	if !ok {
		logLevel = "info"
	}
	splitType, ok := config["split_type"]
	if !ok {
		if splitType != "time" && splitType != "size" {
			splitType = "time"
		}
	}
	SplitIntervalTimeStr, ok := config["Split_interval_time"]
	if !ok {
		SplitIntervalTimeStr = "60"
	}
	splitSizeStr, ok := config["split_size"]
	if !ok {
		splitSizeStr = "100"
	}
	splitIntervalTime, err := strconv.ParseInt(SplitIntervalTimeStr, 10, 64)
	if err != nil {
		splitIntervalTime = 60
	}
	splitSize, err := strconv.ParseInt(splitSizeStr, 10, 64)
	if err != nil {
		splitSize = 100
	}

	logInt := GetLevelInt(logLevel)
	log = &FileLogger{
		LogPath:           LogPath,
		LogName:           logName,
		LogLevel:          logInt,
		LogDataChan:       make(chan *LogData, 5000),
		SplitType:         splitType,
		SplitSize:         splitSize,
		SplitIntervalTime: splitIntervalTime,
		SplitLastTime:     time.Now().Unix(),
	}
	log.Init()
	return log, nil
}
