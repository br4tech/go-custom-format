package port

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"gorm.io/gorm"
)

type (
	IPostgresAdapter interface {
		GetDb() *gorm.DB
	}

	IProductAdapter interface {
		FindBySku(sku string) (*domain.Product, error)
		Create(product *domain.Product) (*domain.Product, error)
	}
)
