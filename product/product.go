package product

import (
	"fmt"
	"math/rand"
)

type Product struct {
	Name          string
	Quantity      uint
	Price         uint
	OriginalPrice uint
}

type ProductSellRequest struct {
	Name     string
	Quantity uint
}

type ProductList []Product

func (p *ProductList) AddProduct(product Product) {
	*p = append(*p, product)
}

func (p *ProductList) RemoveProduct(i int) {
	*p = append((*p)[:i], (*p)[i+1:]...)
}

func NewProduct(name string, price, quantity uint) Product {
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func getProductInfo() ProductSellRequest {
	var (
		productName string
		quantity    uint
	)
	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product Quantity: ")
	fmt.Scan(&quantity)
	fmt.Println()

	return ProductSellRequest{
		Name:     productName,
		Quantity: quantity,
	}
}

func GenerateProductPrice(min, max int) uint {
	return uint(rand.Intn(max-min) + min)
}