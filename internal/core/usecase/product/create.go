package product

import (
	"github.com/br4tech/go-custom-format/internal/core/domain"
	"github.com/br4tech/go-custom-format/internal/core/port"
	"github.com/br4tech/go-custom-format/internal/dto"
	validator "github.com/br4tech/go-custom-format/pkg"
)

type CreateUseCase struct {
	productAdapter port.IProductAdapter
}

func NewCreateUseCase(productAdapter port.IProductAdapter) port.ICreateUseCase {
	return &CreateUseCase{
		productAdapter: productAdapter,
	}
}

func (usecase *CreateUseCase) Execute(productDTO *dto.Product) (*domain.Product, error) {
	product := domain.NewProduct(
		productDTO.Sku,
		productDTO.Name,
		productDTO.Price,
		productDTO.StockLevel,
	)

	if err := validator.ValidateStruct(product); err != nil {
		return nil, err
	}

	createProduct, err := usecase.productAdapter.Create(product)
	if err != nil {
		return nil, err
	}
	return createProduct, nil
}
