package utils

import "fmt"

func init() {
	fmt.Println("utils initialized - [utils.go]")
}

func IsEven(no int) bool {
	return no%2 == 0
}
