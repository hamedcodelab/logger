package logger

import (
	"fmt"
	"github.com/hamedcodelab/configer"
	"github.com/hamedcodelab/logger/customCore"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Fields map[string]interface{}

type logger struct {
	config Config
	logger *zap.Logger
}

func New(register configer.Register) Logger {
	conf := Config{}
	err := register.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    zapcore.ISO8601TimeEncoder,

		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapLevel := zap.NewAtomicLevelAt(getLevel(conf.Level))
	newCore := customCore.NewCore(
		selectEncoder(conf.Encoding, encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapLevel,
	)

	l := zap.New(newCore)

	return &logger{
		config: conf,
		logger: l,
	}
}

func (l *logger) WithFields(fields map[string]interface{}) *logger {
	fld := make([]zapcore.Field, 0)
	for k, field := range fields {
		switch field.(type) {
		case string:
			fld = append(fld, zapcore.Field{
				Key:    k,
				Type:   zapcore.StringType,
				String: field.(string),
			})
		case int:
			fld = append(fld, zapcore.Field{
				Key:     k,
				Type:    zapcore.Int64Type,
				Integer: int64(field.(int)),
			})
		case bool:
			fld = append(fld, zapcore.Field{
				Key:       k,
				Type:      zapcore.BoolType,
				Interface: field.(bool),
			})
		case error:
			fld = append(fld, zapcore.Field{
				Key:       k,
				Type:      zapcore.ErrorType,
				Interface: field.(error),
			})
		default:
			fld = append(fld, zapcore.Field{
				Key:       k,
				Type:      zapcore.ReflectType,
				Interface: field,
			})
		}
	}

	newLogger := l.logger.With(fld...)
	n := &logger{
		config: l.config,
		logger: newLogger,
	}
	return n
}

func (l *logger) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(fmt.Sprintf(msg, fields...))
}

func (l *logger) Info(msg string, fields ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, fields...))
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(fmt.Sprintf(msg, fields...))
}

func (l *logger) Error(msg string, fields ...interface{}) {
	l.logger.Error(fmt.Sprintf(msg, fields...))
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(msg, fields...))
}

func (l *logger) Panic(msg string, fields ...interface{}) {
	l.logger.Panic(fmt.Sprintf(msg, fields...))
}

func selectEncoder(enc string, encConfig zapcore.EncoderConfig) zapcore.Encoder {
	switch enc {
	case "json":
		encConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encConfig)
	case "console":
		encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encConfig)
	default:
		return zapcore.NewConsoleEncoder(encConfig)
	}
}

func getLevel(lvl string) zapcore.Level {
	switch lvl {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
