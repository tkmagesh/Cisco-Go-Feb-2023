package main

import (
	"fmt"

	"github.com/fatih/color"

	calc "github.com/tkmagesh/cisco-go-feb-2023/10-modularity-demo/calculator"
	"github.com/tkmagesh/cisco-go-feb-2023/10-modularity-demo/calculator/utils"
)

func main() {
	fmt.Println("app executed")
	color.Red("app executed")
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Operation Count : ", calc.OpCount())
	fmt.Println("is 21 even ?:", utils.IsEven(21))
}
