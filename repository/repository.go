package repository

import "package_module/product"

type Repository struct {
	Products product.ProductList
}

func NewRepository(products product.ProductList) Repository {
	return Repository{
		Products: products,
	}
}

func (r Repository) Search(productName string) (product.Product, bool) {
	for _, product := range r.Products {
		if product.Name == productName {
			return product, true
		}
	}

	return product.Product{}, false
}

func (r *Repository) TakeProduct(productName string, quantity uint) {
	for i, product := range r.Products {
		if product.Name == productName {
			r.Products[i].Quantity -= quantity
			if r.Products[i].Quantity == 0 {
				r.Products.RemoveProduct(i)
			}
			return
		}
	}
}
