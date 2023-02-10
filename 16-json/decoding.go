package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Units    int     `json:"units"`
	Category string  `json:"category"`
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	var products []Product
	file, err := os.Open("products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		log.Fatalln(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}
