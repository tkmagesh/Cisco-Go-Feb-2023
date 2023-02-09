package calculator

import "fmt"

var op_count map[string]int

func init() {
	op_count = make(map[string]int)
}

func init() {
	fmt.Println("calculator initialized - [calc.go] 1")
}

func init() {
	fmt.Println("calculator initialized - [calc.go] 2")
}

func OpCount() map[string]int {
	return op_count
}
