package sql

import (
	"logger_use/logger"
	"time"
)

func Run() {
	for {
		logger.Error("user server is running Error")
		logger.Warn("user server is running Warn")
		logger.Info("user server is running Info")
		//fmt.Println("是否卡顿")
		time.Sleep(1 * time.Second)
	}
}
