package products

import (
	"context"
	"fmt"
)

var list []Product

func init() {
	list = []Product{
		{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"},
		{Id: 101, Name: "Pencil", Cost: 5, Category: "Stationary"},
		{Id: 102, Name: "Marker", Cost: 50, Category: "Stationary"},
		{Id: 103, Name: "Mouse", Cost: 250, Category: "Electronics"},
	}
}

type ProductsService struct {
}

func NewProductsService() *ProductsService {
	return &ProductsService{}
}

func (ps *ProductsService) GetAll(ctx context.Context) []Product {
	fmt.Println("[products.service - GetAll()] req id :", ctx.Value("request-id"))
	return list[:]
}

func (ps *ProductsService) AddNew(ctx context.Context, product Product) {
	fmt.Println("[products.service - AddNew()] req id :", ctx.Value("request-id"))
	list = append(list, product)
}

func (ps *ProductsService) GetOne(id int) *Product {
	for _, p := range list {
		if p.Id == id {
			return &p
		}
	}
	return nil
}
