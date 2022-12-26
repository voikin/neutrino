package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/voikin/neutrino/models"
)

type jsonMap map[string]interface{}

func (h *Handler) saveTgUser(e echo.Context) error {
	input := &models.TgUser{}

	err := e.Bind(input)
	if err != nil {
		return nil
	}

	id, err := h.repo.TgUserRepository.SaveTgUser(input)
	if err != nil {
		return nil
	}

	return e.JSON(http.StatusOK, jsonMap{
		"message": "successful",
		"id":      id,
	})
}
