package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		fmt.Fprintln(w, "All the products will be served")
	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func main() {
	app := &App{}
	http.ListenAndServe(":8080", app)
}
