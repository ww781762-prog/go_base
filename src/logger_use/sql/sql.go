package sql

import (
	"logger_use/logger"
	"time"
)

func Run() {
	for {
		logger.Error("user server is running")
		time.Sleep(1 * time.Second)
	}
}
