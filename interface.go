package logger

type Logger interface {
	WithFields(fields map[string]interface{}) *logger
	Debug(msg string, fields ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Panic(format string, args ...interface{})
}
