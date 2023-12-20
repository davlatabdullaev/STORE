
package main

import (
	"fmt"
	_"package_module/dealer"
	"package_module/product"
	"package_module/repository"
	"package_module/store"
)

func main() {
	const (
		StartShopCmd = iota + 1
		FinishShopCmd
	)

	repo := repository.NewRepository(product.ProductList{})
	store := store.NewStore(repo)

	for {
		var cmd int
		fmt.Print(`
      Enter command:
      1 - Start Shopping
      2 - Stop Shopping
`)
		fmt.Scan(&cmd)

		switch cmd {
		case StartShopCmd:
			store.PrintStats()
			store.StartSell()
		case FinishShopCmd:
			return
		default:
			fmt.Println("There is not such kind of command!!!")
		}
	}
}
