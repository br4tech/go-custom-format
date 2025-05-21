package model

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Sku        string  `gorm:"column:sku,unique;not null"`
	Name       string  `gorm:"column:name,unique;not null"`
	Price      float64 `gorm:"column:price,unique;not null"`
	StockLevel int     `gorm:"column:stock_level,unique;not null"`
}

func (model Product) ToDomain() *domain.Product {
	return &domain.Product{
		Sku:        model.Sku,
		Name:       model.Name,
		Price:      model.Price,
		StockLevel: model.StockLevel,
	}
}

func (model *Product) FromDomain(domain *domain.Product) {
	model.Sku = domain.Sku
	model.Name = domain.Name
	model.Price = domain.Price
	model.StockLevel = domain.StockLevel
}
