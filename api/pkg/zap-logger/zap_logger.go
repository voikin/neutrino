package zaplogger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(conf zap.Config) *ZapLogger {
	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{
		logger: logger,
	}
}

func (zl *ZapLogger) LoggerMiddleware() echo.MiddlewareFunc {
	lm := middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogHost:   true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			zl.logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	})
	
	return lm
}