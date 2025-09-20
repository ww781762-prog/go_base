package main

import (
	"logger_use/logger"
	"logger_use/sql"
)

// 日志库，每个项目全局只有一个，我们应该在，日志包里面初始化一个全局的对象，后面只使用的时候之际调用就可以了
//var log logger.LoggerInterface
//
//func initLogger(logPath string, logName string, level int) {
//	log = logger.NewFileLogger(logPath, logName, level)
//	log.Debug("init logger success")
//}
//
//func Run() {
//	for {
//		log.Debug("user server is running")
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func main() {
//	initLogger("/Users/wallace/GolandProjects/go_base/src/logger_use/logs", "test", logger.LevelDebug)
//	Run()
//	return
//}

// 封装后的日志库，全局变量封装到了 logger 包里面，初始化后，直接使用就可以了
func initLog() {
	config := map[string]string{
		"log_path":  "/Users/wallace/GolandProjects/go_base/src/logger_use/logs",
		"log_name":  "logger_use",
		"log_level": "debug",
	}
	err := logger.InitLogger("file", config)
	if err != nil {
		panic("fail to init logger")
	}
}

//	func Run() {
//		for {
//			// 全局变量封装在
//			logger.Fatal("user server is running")
//			time.Sleep(1 * time.Second)
//		}
//	}

func main() {
	initLog()
	//Run()
	sql.Run()
}
