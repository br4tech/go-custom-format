package dto

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"github.com/go-playground/validator"
)

type Product struct {
	Sku        string  `json:"sku" binding:"required,skuFormat"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	StockLevel int     `json:"stock_level"`
}

func skuFormatValidator(fl validator.FieldLevel) bool {
	sku := fl.Field().String()

	if len(sku) != 9 {
		return false
	}

	for i, char := range sku {
		switch i {
		case 0, 1, 2:
			if !('A' <= char && char <= 'Z') {
				return false
			}
		case 3:
			if char != '-' {
				return false
			}
		default:
			if !('0' <= char && char <= '9') {
				return false
			}
		}
	}
	return true
}

func init() {
	validate := validator.New()
	validate.RegisterValidation("skuformat", skuFormatValidator)
}

func (dto Product) ToDomain() *domain.Product {
	return &domain.Product{
		Sku:        dto.Sku,
		Name:       dto.Name,
		Price:      dto.Price,
		StockLevel: dto.StockLevel,
	}
}

func (dto *Product) FromDomain(domain *domain.Product) {
	dto.Sku = domain.Sku
	dto.Name = domain.Name
	dto.Price = domain.Price
	dto.StockLevel = domain.StockLevel
}
