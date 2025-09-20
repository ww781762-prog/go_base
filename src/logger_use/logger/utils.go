package logger

import "runtime"

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		// 文件名称
		fileName = file
		// 函数名称
		funcName = runtime.FuncForPC(pc).Name()
		//   行号
		lineNo = line
	}
	return
}
