package port

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"github.com/br4tech/go-custom-format/internal/dto"
)

type (
	ICreateUseCase interface {
		Execute(productDTO *dto.Product) (*domain.Product, error)
	}
	IFindBySkuUseCase interface {
		Execute(sku string) (*domain.Product, error)
	}
)
