package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

func main() {
	cisco := &Organization{Name: "CISCO", City: "Bangalore"}
	e1 := Employee{Id: 100, Name: "Magesh", Salary: 10000, Org: cisco}
	e2 := Employee{Id: 101, Name: "Suresh", Salary: 10000, Org: cisco}
	fmt.Println(e1)
	fmt.Println(e2)

	fmt.Println("After changing the Org city")
	cisco.City = "Pune"
	fmt.Println(e1.Org)
	fmt.Println(e2.Org)

}
