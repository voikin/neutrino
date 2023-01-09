package logger

import (
	"github.com/labstack/echo/v4"
	zaplogger "github.com/voikin/neutrino/pkg/logger/zap-logger"
	"go.uber.org/zap"
)

type Logger struct {
	MVLogger
}

type MVLogger interface {
	LoggerMiddleware() echo.MiddlewareFunc
}

func NewLogger(conf zap.Config) *Logger {
	return &Logger{
		MVLogger: zaplogger.NewZapLogger(conf),
	}
}
