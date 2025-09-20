package loglib

type Logger interface {
	WriteDebug(str string)
	WriteInfo(str string)
	WriteWarn(str string)
}
