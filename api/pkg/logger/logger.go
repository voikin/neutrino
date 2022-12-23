package logger

import (
	zaplogger "github.com/dazai404/neutrino/pkg/logger/zap-logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Logger struct {
	MVLogger
}

type MVLogger interface{
	LoggerMiddleware() echo.MiddlewareFunc
}

func NewLogger(conf zap.Config) *Logger {
	return &Logger{
		MVLogger: zaplogger.NewZapLogger(conf),
	}
}


