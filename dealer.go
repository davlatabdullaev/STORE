package main

type Dealer struct{}

func (d Dealer) ProvideProduct(budget uint, productName string, quantity uint) (Product, bool) {
	OriginalPrice := generateProductPrice(1, 10) * 1000
	temp := 0
	if budget < OriginalPrice*uint(quantity) {
		temp = int(budget) / int(OriginalPrice)

		if temp < 1 {
			
			return Product{}, false
		}
		} else {
		temp = int(quantity)
	}

	return Product{
		Name:          productName,
		Price:         OriginalPrice * 6 / 5,
		Quantity:      uint(temp),
		OriginalPrice: OriginalPrice,
	}, true
}
