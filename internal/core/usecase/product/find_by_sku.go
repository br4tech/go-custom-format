package product

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"github.com/br4tech/go-custom-format/internal/core/port"
)

type FindBySkuUseCase struct {
	productAdapter port.IProductAdapter
}

func NewFindBySkuUseCase(productAdapter port.IProductAdapter) port.IFindBySkuUseCase {
	return &FindBySkuUseCase{
		productAdapter: productAdapter,
	}
}

func (usecase *FindBySkuUseCase) Execute(sku string) (*domain.Product, error) {
	product, err := usecase.productAdapter.FindBySku(sku)
	if err != nil {
		return nil, err
	}
	return product, nil
}
