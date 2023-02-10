package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

var products = []Product{
	Product{105, "Pen", 5, 50, "Stationary"},
	Product{107, "Pencil", 2, 100, "Stationary"},
	Product{103, "Marker", 50, 20, "Utencil"},
	Product{102, "Stove", 5000, 5, "Utencil"},
	Product{101, "Kettle", 2500, 10, "Utencil"},
	Product{104, "Scribble Pad", 20, 20, "Stationary"},
	Product{109, "Golden Pen", 2000, 20, "Stationary"},
}

//handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/customers", CustomersHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
