package main

import "fmt"

func main() {
	emp := struct {
		Id     int
		Name   string
		Salary float32
	}{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Println(emp)
	PrintEmployee(emp)

	signalObj := struct{}{}
	fmt.Println(signalObj)
}

func PrintEmployee(e struct {
	Id     int
	Name   string
	Salary float32
}) {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", e.Id, e.Name, e.Salary)
}
