package main

import (
	"fmt"
)

func main() {
	// var x interface{} // => before go1.18
	var x any
	// x = 100
	// x = "Proident ex velit dolor ut velit. Occaecat anim in nostrud veniam. Quis ipsum sint dolor voluptate in dolor eiusmod dolor anim in dolor."
	// x = true
	// x = 87.87
	// x = struct{}{}
	x = []int{}

	// x = 1000
	// x = "Labore minim irure exercitation ut excepteur velit ipsum. Mollit pariatur minim adipisicing reprehenderit proident exercitation Lorem. Cupidatat magna dolore incididunt sunt dolore occaecat laborum veniam pariatur in commodo quis. Laboris consequat magna elit eiusmod quis tempor. Exercitation nostrud consectetur Lorem amet cupidatat consequat magna est nisi commodo non magna. Occaecat aliqua enim tempor exercitation ut id nulla exercitation."
	// fmt.Println(x.(int) + 2000)
	if val, ok := x.(int); ok {
		fmt.Println(val + 2000)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 2000 =", val+2000)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float, 99.9% of x =", val*(99.9/100))
	case struct{}:
		fmt.Println("x is an empty struct")
	default:
		fmt.Println("unknown type")
	}
}
