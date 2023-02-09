package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	//type inference
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using indexer")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	sort(&nos)
	fmt.Println("After sorting, nos =", nos)

	productsList1 := [3]string{"Pen", "Pencil", "Marker"}
	productsList2 := [3]string{"Pen", "Pencil", "Marker"}
	fmt.Printf("%p, %p\n", &productsList1, &productsList2)
	fmt.Println(productsList1 == productsList2)
}

func sort(list *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			/*
				if (*list)[i] > (*list)[j] {
					(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
				}
			*/
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
