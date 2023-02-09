package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

//fmt.Stringer implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

//fmt.Stringer implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func main() {
	grapes := PerishableProduct{
		Product: Product{100, "Grapes", 10},
		Expiry:  "2 Days",
	}

	fmt.Println(grapes)
	fmt.Println("After applying discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes)

}
