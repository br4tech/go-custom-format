package domain

type Product struct {
	Sku        string
	Name       string
	Price      float64
	StockLevel int
}

func NewProduct(sku string, name string, price float64, stock_level int) *Product {
	return &Product{
		Sku:        sku,
		Name:       name,
		Price:      price,
		StockLevel: stock_level,
	}
}
