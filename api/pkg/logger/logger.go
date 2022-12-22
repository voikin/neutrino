package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(conf zap.Config) *Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Logger{
		logger: l,
	}
}

func (l *Logger) LoggerMiddleware() echo.MiddlewareFunc {
	lm := middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogHost:   true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	})
	
	return lm
}
