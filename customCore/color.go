package customCore

import "go.uber.org/zap/zapcore"

func getZapLevelColor(level zapcore.Level) int {
	switch level {
	case zapcore.DebugLevel:
		return 36 // Cyan
	case zapcore.InfoLevel:
		return 32 // Green
	case zapcore.WarnLevel:
		return 33 // Yellow
	case zapcore.ErrorLevel:
		return 31 // Red
	case zapcore.FatalLevel:
		return 91 // Bright Red
	case zapcore.PanicLevel:
		return 35 // Magenta
	default:
		return 0 // Default (no color)
	}
}
