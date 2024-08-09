package service

import (
	"avito_bootcamp/internal/entity"
	"avito_bootcamp/internal/validation"
	"avito_bootcamp/pkg/sender"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func (s *Services) GetHouse(c echo.Context) error {
	tokenString, err := validation.GetToken(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	claims, err := validation.CheckToken(tokenString)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	param := c.Param("id")
	num, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request number")
	}
	flats, err := s.Storage.GetHouse(context.Background(), num, claims.UserType)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	for _, flat := range flats {
		c.JSON(http.StatusOK, echo.Map{"number": flat.Number, "house_id": flat.HouseID, "price": flat.Price, "rooms": flat.Rooms, "status": flat.Status})
	}
	return nil
}

func (s *Services) CreateHouse(c echo.Context) error {
	tokenString, err := validation.GetToken(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	err = validation.CheckModerator(tokenString)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	var house entity.House
	if err = c.Bind(&house); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	err = s.Storage.CreateHouse(context.Background(), house)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"id": house.ID, "address": house.Address, "year_built": house.YearBuilt, "developer": house.Developer, "created_at": house.CreatedAt})
}

func (s *Services) SubscribeHouse(c echo.Context) error {
	tokenString, err := validation.GetToken(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	_, err = validation.CheckToken(tokenString)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	var user entity.User
	if err = c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	go func() {
		send := sender.New()
		for {
			time.Sleep(30 * time.Second)
			err = send.SendEmail(context.Background(), user.Email, "example of sending message")
			if err != nil {
				return
			}
		}
	}()
	return nil
}
