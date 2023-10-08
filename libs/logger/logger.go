package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type log struct{}

func New() Logger {
	logger_cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			// message
			MessageKey: "MESSAGE",
			// callstack
			CallerKey:    "CALL FROM",
			EncodeCaller: zapcore.ShortCallerEncoder,
			// time
			TimeKey:    "TIME",
			EncodeTime: zapcore.TimeEncoderOfLayout("__15:04:05__"),
			// name of the field already says what it is :)
			LineEnding: "\t__EOF__\n",
		},
	}

	logger, err := logger_cfg.Build()
	defer logger.Sync()

	if err != nil {
		panic(err)
	}

	surgared_logger := logger.Sugar()
	return surgared_logger
}
