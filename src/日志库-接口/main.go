package main

import "loglib/loglib"

func main() {
	// 调用起来比较麻烦，修改了写入的地方整体都需要修改
	//filelog := loglib.NewFileLog()
	//filelog.FileInfLog("hello world")
	//filelog.FileDebugLog("ww")
	//
	//consolLog := loglib.NewConsolLog()
	//consolLog.ConsolInfLog("hello world")
	//consolLog.ConsolDebugLog("ww")

	// 通过接口实现
	// 指用修改了实例化就可以了
	//logs := loglib.NewFileLog("File")
	log := loglib.NewConsolLog("File")
	// 下面的都不需要修改
	log.WriteDebug("hello")
	log.WriteInfo("hello")
	log.WriteWarn("hello")
}
