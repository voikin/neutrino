package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/voikin/neutrino/models"
)

type jsonMap map[string]interface{}

func (h *Handler) saveTgUser(e echo.Context) error {
	input := &models.TgUser{}
	err := e.Bind(input)
	if err != nil {
		return err
	}
	if input.UserId == nil {
		return echo.ErrBadRequest
	}
	_, err = h.repo.GetTgUser(*input.UserId)
	if err == nil {
		err = h.repo.TgUserRepository.UpdateTgUser(input)
		fmt.Println(err)
		if err != nil {
			return err
		}
		return e.JSON(http.StatusOK, jsonMap{
			"message": "successfully updated",
		})
	}
	id, err := h.repo.TgUserRepository.SaveTgUser(input)
	if err != nil {
		return nil
	}
	return e.JSON(http.StatusCreated, jsonMap{
		"message": "successfully created",
		"id":      id,
	})
}
