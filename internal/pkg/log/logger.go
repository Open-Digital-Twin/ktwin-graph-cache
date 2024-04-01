package log

import "go.uber.org/zap"

func NewLogger() Logger {
	zapLogger, _ := zap.NewProduction()
	return &logger{
		lg: zapLogger,
	}
}

type Logger interface {
	Info(message string)
	Error(message string)
}

type logger struct {
	lg *zap.Logger
}

func (l *logger) Info(message string) {
	l.lg.Info(message)
}

func (l *logger) Error(message string) {
	l.lg.Error(message)
}
