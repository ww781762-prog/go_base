package logger

import "testing"

func TestFileLogger(t *testing.T) {
	config := map[string]string{
		"log_path":  "/Users/wallace/GolandProjects/go_base/src/logger_use/logs",
		"log_name":  "logger_use",
		"log_level": "warn",
	}
	err := InitLogger("file", config)
	if err != nil {
		panic("fail to init logger")
	}
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

//func TestConsoleLogger(t *testing.T) {
//	//log := NewConsoleLogger(LevelDebug)
//	log.Debug("debug")
//	log.Info("info")
//	log.Warn("warn")
//	log.Error("error")
//}
