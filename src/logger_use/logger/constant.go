package logger

const (
	LevelFatal = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelTrace
	LevelDebug
)

func GetLevelText(level int) string {
	switch level {
	case LevelFatal:
		return "fatal"
	case LevelError:
		return "error"
	case LevelWarn:
		return "warn"
	case LevelInfo:
		return "info"
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	default:
		return "info"
	}
}
func GetLevelInt(level string) int {
	switch level {
	case "fatal":
		return LevelFatal
	case "error":
		return LevelError
	case "warn":
		return LevelWarn
	case "info":
		return LevelInfo
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	default:
		return LevelInfo
	}
}
