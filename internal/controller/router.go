package controller

import (
	"avito_bootcamp/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.Use(middleware.Logger())
	handler.Use(middleware.Recover())

	handler.GET("/dummyLogin", services.DummyLogin)
	handler.POST("/login", services.Login)
	handler.POST("/register", services.Register)
	handler.GET("/house/:id", services.GetHouse)
	handler.POST("/flat/create", services.CreateFlat)
	handler.POST("/house/:id/subscribe", services.SubscribeHouse)
	handler.POST("/house/create", services.CreateHouse)
	handler.POST("/flat/update", services.UpdateFlat)
}
