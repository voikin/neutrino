package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	input.UpdatedAt = time.Now().UTC()
	tgUser, err := h.repo.GetTgUser(*input.UserId)
	if err == nil {
		input.CreatedAt = tgUser.CreatedAt
		err = h.repo.TgUserRepository.UpdateTgUser(input)
		fmt.Println(err)
		if err != nil {
			return err
		}
		return e.JSON(http.StatusNoContent, nil)
	}
	input.CreatedAt = time.Now().UTC()
	err = h.repo.TgUserRepository.SaveTgUser(input)
	if err != nil {
		return nil
	}
	return e.JSON(http.StatusCreated, jsonMap{
		"message": "successfully created",
	})
}

func (h *Handler) getTgUser(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	user, err := h.repo.TgUserRepository.GetTgUser(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, user)
}

func (h *Handler) deleteTgUser(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = h.repo.TgUserRepository.DeleteTgUser(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusNoContent, nil)
}
