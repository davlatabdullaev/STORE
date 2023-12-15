package main

func main() {
	var (
		productList = ProductList{
			{Name: "Non", Price: 4000, Quantity: 10},
			{Name: "Cola", Price: 13000, Quantity: 15},
			{"Nestle", 3000, 20},
		}

		repository = Repository{}
		store      = Store{}
	)

	newRepository := repository.NewRepository(productList)

	newStore := store.NewStore(newRepository)

	newStore.Sell()
}
