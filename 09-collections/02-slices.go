package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4}
	nos := []int{3, 1, 4}
	fmt.Println(nos)

	fmt.Println("Iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos = append(nos, 10)
	fmt.Println(nos)
	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	myNos := []int{3, 1, 4, 2, 5}
	sort(myNos) //=> ?
	fmt.Println("After sorting : myNos =", myNos)

	//slicing
	/*
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end
		[:hi] => from 0 to hi-1
	*/
	fmt.Println("nos[2:4] = ", nos[2:4])
	fmt.Println("nos[2:] = ", nos[2:])
	fmt.Println("nos[:4] = ", nos[:4])

	subNos := nos[2:4]
	subNos[0] = 10000
	fmt.Println("subNos :", subNos)
	fmt.Println("nos :", nos)
	fmt.Println("len(subNos) :", len(subNos))
	fmt.Println("subNos[5] :", subNos[5])

}

func sort(list []int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
