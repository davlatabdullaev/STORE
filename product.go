package main

import "fmt"

type Product struct {
	Name     string
	Price    uint
	Quantity uint
}
type ProductSellRequest struct {
	Name     string
	Quantity uint
}

type ProductList []Product

func (p Product) NewProduct(name string, price, quantity uint) Product {
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
