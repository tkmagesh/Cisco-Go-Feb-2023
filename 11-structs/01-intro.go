package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func NewEmployee(id int, name string, salary float32) *Employee {
	return &Employee{
		Id:     id,
		Name:   name,
		Salary: salary,
	}
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000}
	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	fmt.Printf("%+v\n", emp)

	//structs are values
	e1 := Employee{200, "Suresh", 10000}
	e2 := Employee{200, "Suresh", 10000}
	fmt.Println(e1 == e2)

	fmt.Println("Before awarding bonus, emp = ", emp)
	AwardBonus(&emp, 2000)
	fmt.Println("After awarding bonus, emp = ", emp)

	// empPtr := &Employee{}
	// empPtr := new(Employee)
	empPtr := NewEmployee(200, "Suresh", 20000)
	fmt.Println(empPtr)
}

func AwardBonus(emp *Employee, bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}
