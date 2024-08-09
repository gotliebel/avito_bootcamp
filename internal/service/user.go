package service

import (
	"avito_bootcamp/internal/entity"
	"avito_bootcamp/internal/validation"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Services) DummyLogin(c echo.Context) error {
	level := c.QueryParam("user_type")
	if _, ok := entity.UserType[level]; !ok {
		return c.String(http.StatusForbidden, "Wrong type of user")
	}
	tokenString, err := validation.CreateToken(entity.User{UserType: level})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong during creation of token")
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
	})
	return c.JSON(http.StatusOK, echo.Map{"message": "dummy login successful", "token": tokenString})
}

func (s *Services) Login(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		fmt.Println(user.Email, user.Password)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	user, err := s.Storage.Login(context.Background(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	tokenString, err := validation.CreateToken(entity.User{Email: user.Email, UserType: user.UserType})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong during creation of token")
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
	})
	return c.JSON(http.StatusOK, echo.Map{"message": "login successful", "token": tokenString})
}

func (s *Services) Register(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		fmt.Println(user.Email, user.Password)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if user.UserType != "client" && user.UserType != "moderator" {
		return c.String(http.StatusUnauthorized, "Wrong type of user")
	}

	err := s.Storage.Register(context.Background(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "OK")
}
