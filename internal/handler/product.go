package handler

import (
	"net/http"

	"github.com/br4tech/go-custom-format/internal/core/port"
	"github.com/br4tech/go-custom-format/internal/dto"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productCreateUseCase    port.ICreateUseCase
	productFindBySkuUseCase port.IFindBySkuUseCase
}

func NewProductHandler(productCreateUseCase port.ICreateUseCase,
	productFindBySkuUseCase port.IFindBySkuUseCase) *ProductHandler {
	return &ProductHandler{
		productCreateUseCase:    productCreateUseCase,
		productFindBySkuUseCase: productFindBySkuUseCase,
	}
}

func (handler *ProductHandler) FindBySku(c echo.Context) error {
	sku := c.Param("sku")

	product, err := handler.productFindBySkuUseCase.Execute(sku)
	if err != nil {
		return HandlerResponse(c, http.StatusNotFound, "Product not found")
	}

	return HandlerResponse(c, http.StatusOK, product)
}

func (handler *ProductHandler) Create(c echo.Context) error {
	reqBody := new(dto.Product)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	product, err := handler.productCreateUseCase.Execute(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to created Product")
	}

	return HandlerResponse(c, http.StatusCreated, product)
}
