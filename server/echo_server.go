package server

import (
	"fmt"

	"github.com/br4tech/go-custom-format/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	app *echo.Echo
	cfg *config.Config
}

func NewechoServer(cfg *config.Config) Server {
	return &EchoServer{
		app: echo.New(),
		cfg: cfg,
	}
}

func (s *EchoServer) Start() {
	s.app.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
		}))

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
