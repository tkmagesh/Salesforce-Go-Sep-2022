package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/* bread := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Bread",
			Cost: 10,
		},
		Expiry: "3 Days",
	} */
	bread := NewPerishableProduct(100, "bread", 10, "3 Days")
	fmt.Println(bread)
	fmt.Println(bread.Product.Cost)
	fmt.Println(bread.Cost)
}
