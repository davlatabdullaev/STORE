package main

import "fmt"

func main() {
const(
	StartShopCmd = iota+1
	FinishShopCmd
)

	repo := NewRepository(productList)
	store := NewStore(repo)
	

	for true{
var cmd int 
fmt.Print(`
      Enter command:
      1 -  Start Shopping
      2 -  Stop Shopping
`)
fmt.Scan(&cmd)

switch cmd{
case StartShopCmd :
	store.printStats()
	store.StartSell()
	store.printStats()
case FinishShopCmd:
	return
}
}
}