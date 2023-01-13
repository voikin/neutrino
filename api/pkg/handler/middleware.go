package handler

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *Handler) authBotMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			token := e.QueryParam("token")

			h := sha256.New()
			h.Write([]byte(os.Getenv("BOT_KEY")))
			hash := fmt.Sprintf("%x", h.Sum(nil))
			fmt.Println(hash)

			if token != string(hash) {
				return echo.ErrUnauthorized
			}

			return next(e)
		}
	}
}
