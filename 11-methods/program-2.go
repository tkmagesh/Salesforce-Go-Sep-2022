package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product *Product) ApplyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v", product.Id, product.Name, product.Cost)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp *PerishableProduct) Format() string {
	//return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Expiry = %q", pp.Id, pp.Name, pp.Cost, pp.Expiry)
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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
	pen := Product{100, "Pen", 10}
	//fmt.Println(pen)
	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	//fmt.Println(pen)
	fmt.Println(pen.Format())

	bread := NewPerishableProduct(101, "Bread", 20, "3 Days")
	//fmt.Println(bread.Cost)
	fmt.Println(bread.Format())
	bread.ApplyDiscount(10)
	fmt.Println(bread.Format())
}
