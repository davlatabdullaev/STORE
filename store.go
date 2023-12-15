package main
 
import "fmt"

type Store struct {
	Repository Repository
	Cash       uint
}

func (s Store) Sell() {
	s.printStats()

	productSellRequest := getProductInfo()

	product, exits := s.Repository.Search(productSellRequest.Name)
	if !exits {
		fmt.Printf("We do not have %s product", productSellRequest.Name)
		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, letf %d", productSellRequest.Name, product.Quantity)
		return
	}
	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)
	s.Cash += productSellRequest.Quantity * product.Price
     s.printStats()
}

func (s Store) NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Cash:       0,
	}
}
func (s Store) printStats() {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("                                         Cash: ", s.Cash)
	for _, product := range s.Repository.Products {
		fmt.Println(product.Name, product.Quantity, product.Price)
		fmt.Println()
	}
	fmt.Println("--------------------------------------------------------")
}

