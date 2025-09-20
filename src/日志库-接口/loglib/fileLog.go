package loglib

import "fmt"

type fileLog struct {
	Name string
}

//func NewFileLog() *fileLog {
//	return &fileLog{}
//}
//
//func (*fileLog) FileDebugLog(str string) {
//	fmt.Println("Debug", str)
//}
//
//func (*fileLog) FileInfLog(str string) {
//	fmt.Println("INFO", str)
//}
//

func NewFileLog(Name string) Logger {
	return &fileLog{Name: Name}
}

func (f *fileLog) WriteWarn(str string) {
	fmt.Printf("%s  Warn: %s\n", f.Name, str)
}

func (f *fileLog) WriteInfo(str string) {
	fmt.Printf("%s  Info: %s\n", f.Name, str)

}

func (f *fileLog) WriteDebug(str string) {
	fmt.Printf("%s  Debug: %s\n", f.Name, str)

}
