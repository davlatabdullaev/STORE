package main

import (
	"fmt"
	"math/rand"
)

var (
	productList = ProductList{
		{Name: "Non", Price: 4000, Quantity: 10, OriginalPrice: 3000},
		{Name: "Cola", Price: 13000, Quantity: 15, OriginalPrice: 10_000},
		{Name: "Nestle", Price: 3000, Quantity: 20, OriginalPrice: 2000},
	}
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

func (p *ProductList) RemoveProduct(i int){
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
func generateProductPrice(min, max int) uint{
return uint(rand.Intn(max-min)+min)
}