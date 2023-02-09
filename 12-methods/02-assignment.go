package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

/* write the Format() and ApplyDiscount() methods for the Product type */

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

//overriding the Product.Format()
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	grapes := PerishableProduct{
		Product: Product{100, "Grapes", 10},
		Expiry:  "2 Days",
	}

	fmt.Println(grapes.Format())
	fmt.Println("After applying discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())

}
