package logger

import "fmt"

var log LoggerInterface

// name ：表示实例化一个 文件日志 还是一个console日志
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		log, err = NewFileLogger(config)
	case "console":
		log, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unknown logger %s", name)
	}
	return err
}

func Debug(format string, args ...interface{}) {
	log.Debug(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	log.Info(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	log.Warn(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	log.Error(fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(fmt.Sprintf(format, args...))
}
