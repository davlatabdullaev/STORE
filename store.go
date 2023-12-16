package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Store struct {
	Repository Repository
	Cash       uint
}

func (s *Store) StartSell() {
	

	productSellRequest := getProductInfo()

	product, exits := s.Repository.Search(productSellRequest.Name)
	if !exits {
		fmt.Printf("We do not have %s product\n", productSellRequest.Name)
		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, letf %d\n", productSellRequest.Name, product.Quantity)
		return
	}
	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)
	s.Cash += productSellRequest.Quantity * product.Price

}

func NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Cash:       0,
	}
}
func (s *Store) printStats() {
f:= tabwriter.NewWriter(os.Stdout, 2, 8, 1, '\t', 0)	
fmt.Fprintln(f, "Name\tQuantity\tPrice\t")
for _, product := range s.Repository.Products{
	fmt.Fprintf(f, "%s\t%d\t%d\n", product.Name, product.Quantity, product.Price)
}
fmt.Fprintf(f, "\t\t\tCash : %d\n", s.Cash)
f.Flush()
}

