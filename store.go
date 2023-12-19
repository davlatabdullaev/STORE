package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Store struct {
	Dealer     Dealer
	Repository Repository
	Budget     uint
	Profit     uint
}

func (s *Store) StartSell() {

	productSellRequest := getProductInfo()

	product, exits := s.Repository.Search(productSellRequest.Name)
	if !exits {
		fmt.Printf("We do not have %s product\n We will bring %s in the next time \n", productSellRequest.Name, productSellRequest.Name)
		
		product, succes := s.Dealer.ProvideProduct(s.Budget, productSellRequest.Name, (productSellRequest.Quantity))
		if !succes{
			fmt.Println("We will buy!!!")
			return
		}
		
		s.Repository.Products.AddProduct(product)
	
	    s.Budget -= product.OriginalPrice*product.Quantity
	
		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, letf %d\n", productSellRequest.Name, product.Quantity)
		return
	}
	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)
	s.Profit += productSellRequest.Quantity * (product.Price-product.OriginalPrice)
	
	s.printStats()
}

func NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Budget:     10_000,
		Profit:     0,
	}
}
func (s *Store) printStats() {
	f := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	fmt.Fprintln(f, "Name\tQuantity\tPrice\tOriginal Price\t")
	for _, product := range s.Repository.Products {
		fmt.Fprintf(f, "%s\t%d\t%d\t%d\n", product.Name, product.Quantity, product.Price, product.OriginalPrice)
	}
	fmt.Fprintf(f, "\t\t\tBudget : %d\n", s.Budget)
	fmt.Fprintf(f, "\t\t\tProfit : %d\n", s.Profit)
	f.Flush()
}
