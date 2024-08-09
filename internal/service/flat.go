package service

import (
	"avito_bootcamp/internal/entity"
	"avito_bootcamp/internal/validation"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func (s *Services) CreateFlat(c echo.Context) error {
	tokenString, err := validation.GetToken(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	_, err = validation.CheckToken(tokenString)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	var flat entity.Flat
	if err = c.Bind(&flat); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request with flat")
	}

	err = s.Storage.CreateFlat(context.Background(), flat)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"number": flat.Number, "house_id": flat.HouseID, "price": flat.Price, "rooms": flat.Rooms, "status": flat.Status})

}

func (s *Services) UpdateFlat(c echo.Context) error {
	tokenString, err := validation.GetToken(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	err = validation.CheckModerator(tokenString)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	var flat entity.Flat
	if err = c.Bind(&flat); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	mux := sync.Mutex{}
	mux.Lock()
	defer mux.Unlock()
	flat.Status = "on moderate"
	err = s.Storage.UpdateFlat(context.Background(), flat)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"number": flat.Number, "house_id": flat.HouseID, "price": flat.Price, "rooms": flat.Rooms, "status": flat.Status})
}
