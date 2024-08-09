package app

import (
	"avito_bootcamp/internal/controller"
	"avito_bootcamp/internal/service"
	"avito_bootcamp/internal/storage"
	"avito_bootcamp/pkg/httpserver"
	"fmt"
	_ "github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	st, err := storage.NewStorage()
	if err != nil {
		return
	}

	defer st.Close()

	services := service.New(st)
	e := echo.New()
	controller.NewRouter(e, services)
	httpServer := httpserver.New(e, httpserver.Port("8081"))
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
