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
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	file, err := os.Create("products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(products); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done")
}
