package store

import (
	"fmt"
	"os"
	"text/tabwriter"
	"package_module/dealer"
	"package_module/repository"
	"package_module/product"
)

type Store struct {
	Dealer     dealer.Dealer
	Repository repository.Repository
	Budget     uint
	Profit     uint
}

func (s *Store) StartSell() {
	productSellRequest := s.getProductInfo()

	product, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have %s product\n We will bring %s in the next time \n", productSellRequest.Name, productSellRequest.Name)

		product, success := s.Dealer.ProvideProduct(s.Budget, productSellRequest.Name, productSellRequest.Quantity)
		if !success {
			fmt.Println("We will buy!!!")
			return
		}

		s.Repository.Products.AddProduct(product)
		s.Budget -= product.OriginalPrice * product.Quantity

		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enough %s product, left %d\n", productSellRequest.Name, product.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)
	s.Profit += productSellRequest.Quantity * (product.Price - product.OriginalPrice)

	s.PrintStats()
}

func (s *Store) getProductInfo() product.ProductSellRequest {
	var (
		productName string
		quantity    uint
	)
	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product Quantity: ")
	fmt.Scan(&quantity)
	fmt.Println()

	return product.ProductSellRequest{
		Name:     productName,
		Quantity: quantity,
	}
}

func NewStore(repository repository.Repository) Store {
	return Store{
		Repository: repository,
		Budget:     10_000,
		Profit:     0,
	}
}

func (s *Store) PrintStats() {
	f := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	fmt.Fprintln(f, "Name\tQuantity\tPrice\tOriginal Price\t")
	for _, product := range s.Repository.Products {
		fmt.Fprintf(f, "%s\t%d\t%d\t%d\n", product.Name, product.Quantity, product.Price, product.OriginalPrice)
	}
	fmt.Fprintf(f, "\t\t\tBudget : %d\n", s.Budget)
	fmt.Fprintf(f, "\t\t\tProfit : %d\n", s.Profit)
	f.Flush()
}