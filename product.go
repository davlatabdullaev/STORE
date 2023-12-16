package main

import "fmt"

var (
	productList = ProductList{
		{Name: "Non", Price: 4000, Quantity: 10},
		{Name: "Cola", Price: 13000, Quantity: 15},
		{Name: "Nestle",Price: 3000,Quantity:  20},
	}
)

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
