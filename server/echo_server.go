package server

import (
	"fmt"

	"github.com/br4tech/go-custom-format/config"
	"github.com/br4tech/go-custom-format/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	app            *echo.Echo
	cfg            *config.Config
	productHandler port.IProductHandler
}

func NewechoServer(cfg *config.Config, productHandler port.IProductHandler) Server {
	return &EchoServer{
		app:            echo.New(),
		cfg:            cfg,
		productHandler: productHandler,
	}
}

func (s *EchoServer) Start() {
	s.app.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
		}))

	s.app.POST("/product", s.productHandler.Create)
	s.app.GET("/product/:sku", s.productHandler.FindBySku)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
