package loglib

import "fmt"

type conSolLog struct {
	Name string
}

//	func NewConsolLog() *conSolLog {
//		return &conSolLog{}
//	}
//
//	func (*conSolLog) ConsolDebugLog(str string) {
//		fmt.Println("Debug", str)
//	}
//
//	func (*conSolLog) ConsolInfLog(str string) {
//		fmt.Println("INFO", str)
//	}
func NewConsolLog(Name string) Logger {
	return &conSolLog{Name: Name}
}

func (c *conSolLog) WriteWarn(str string) {
	fmt.Printf("%s  Warn: %s\n", c.Name, str)
}

func (c *conSolLog) WriteInfo(str string) {
	fmt.Printf("%s  Info: %s\n", c.Name, str)
}

func (c *conSolLog) WriteDebug(str string) {
	fmt.Printf("%s  Debug: %s\n", c.Name, str)
}
