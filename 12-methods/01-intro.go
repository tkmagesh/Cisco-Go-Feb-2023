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

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) AwardBonus(bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}

func main() {
	emp := Employee{100, "Magesh", 10000}
	// fmt.Println(Format(emp))
	fmt.Println(emp.Format())

	fmt.Println("After awarding bonus....")
	// AwardBonus(&emp, 2000)
	// (&emp).AwardBonus(2000)
	emp.AwardBonus(2000)
	fmt.Println(emp.Format())
}
