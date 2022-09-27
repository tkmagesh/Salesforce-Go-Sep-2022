package main

import "fmt"

func main() {
	//var productRanks map[string]int
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 4
	*/
	//productRanks := map[string]int{"pen": 4, "pencil": 1, "marker": 2}
	productRanks := map[string]int{
		"pen":    4,
		"pencil": 1,
		"marker": 2,
	}
	fmt.Println(productRanks)

	//accessing the value using the key
	fmt.Printf("Rank of marker is %d\n", productRanks["marker"])
	fmt.Printf("Rank of notebook is %d\n", productRanks["notebook"])

	fmt.Println("Check the existence of a key")
	keyToCheck := "notebook"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Product with key %q does not exist\n", keyToCheck)
	}

	fmt.Println("len(productRanks) = ", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Deleting pen")
	delete(productRanks, "pen")
	fmt.Println(productRanks)
}
