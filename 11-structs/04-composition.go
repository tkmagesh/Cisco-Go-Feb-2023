package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	// Dummy
	Product
	Id     string
	Expiry string
}

func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 10}, "2 Days"}
	grapes := PerishableProduct{
		Product: Product{100, "Grapes", 10},
		Id:      "AG-101",
		Expiry:  "2 Days",
	}
	fmt.Println(grapes)

	//accessing the fields
	// fmt.Println(grapes.Product.Id, grapes.Product.Name, grapes.Product.Cost, grapes.Expiry)
	fmt.Println(grapes.Id, grapes.Name, grapes.Cost, grapes.Expiry)
}
