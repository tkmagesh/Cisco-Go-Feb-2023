package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	// handlers map[string]func(http.ResponseWriter, *http.Request)
	handlers map[string]http.HandlerFunc
}

func NewApp() *App {
	return &App{
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (app *App) Register(pattern string, handler http.HandlerFunc) {
	app.handlers[pattern] = handler
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	if handler, found := app.handlers[r.URL.Path]; found {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the products will be served")
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	app := NewApp()
	app.Register("/", IndexHandler)
	app.Register("/products", ProductsHandler)
	app.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", app)
}
