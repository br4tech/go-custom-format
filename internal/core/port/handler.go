package port

import "github.com/labstack/echo/v4"

type (
	IProductHandler interface {
		FindBySku(c echo.Context) error
		Create(c echo.Context) error
	}
)
