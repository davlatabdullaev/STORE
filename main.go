package main

func main() {
var (
		repository = Repository{}
		store      = Store{}
	)

	newRepository := repository.NewRepository(productList)

	newStore := store.NewStore(newRepository)

	newStore.Sell()
}
