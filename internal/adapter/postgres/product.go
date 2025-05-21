package postgres

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"github.com/br4tech/go-custom-format/internal/core/port"
	"github.com/br4tech/go-custom-format/internal/model"
	"gorm.io/gorm"
)

type ProductAdapter struct {
	Db *gorm.DB
}

func NewProductAdapter(db *gorm.DB) port.IProductAdapter {
	return &ProductAdapter{Db: db}
}

func (adapter *ProductAdapter) FindBySku(sku string) (*domain.Product, error) {
	var product *model.Product

	if err := adapter.Db.Where("sku=?", sku).First(&product).Error; err != nil {
		return nil, err
	}

	return product.ToDomain(), nil
}

func (adapter *ProductAdapter) Create(product *domain.Product) (*domain.Product, error) {
	productModel := new(model.Product)
	productModel.FromDomain(product)

	if err := adapter.Db.Create(productModel).Error; err != nil {
		return nil, err
	}

	return productModel.ToDomain(), nil
}
