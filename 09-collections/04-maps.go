package main

import "fmt"

func main() {
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{}
	// productRanks := make(map[string]int)
	// productRanks := map[string]int{"marker": 3, "pencil": 1}
	productRanks := map[string]int{
		"marker": 3,
		"pencil": 1,
	}

	//adding a new key/value pair
	productRanks["pen"] = 4
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iterating a map using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	//check for the existance of a key
	keyToCheck := "notepad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %s is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] does not exist\n", keyToCheck)
	}

	//removing a key/value pair
	keyToRemove := "pen"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
