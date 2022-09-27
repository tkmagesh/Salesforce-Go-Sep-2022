package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//appending value
	nos = append(nos, 7)
	fmt.Println(nos)

	//appending more values
	nos = append(nos, 6, 8, 9)
	fmt.Println(nos)

	//append values from another slice
	tens := []int{10, 20, 30}
	nos = append(nos, tens...)
	fmt.Println(nos)

	fmt.Println("Iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		slicing
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end of the slice
		[:hi] => from 0 to hi-1
	*/

	fmt.Println("nos[3:6] =", nos[3:6])
	fmt.Println("nos[3:] =", nos[3:])
	fmt.Println("nos[:6] =", nos[:6])

	/*
		//copy of the pointer (both the slices point to the same underlying array)
		newNos := nos
		newNos[0] = 1000
		fmt.Println(newNos, nos)
	*/

	/*
		nosSubset := nos[3:6]
		fmt.Println(len(nosSubset), nosSubset)
		nosSubset[0] = 1000
		fmt.Println("nosSubset = ", nosSubset)
		fmt.Println("nos = ", nos)
		nos = append(nos, 999)
		fmt.Println("nosSubset = ", nosSubset)
		fmt.Println("nos = ", nos)
	*/

	//creating a copy of the slice with data
	copyOfNos := make([]int, len(nos), cap(nos))
	copy(copyOfNos, nos)
	copyOfNos[0] = 1000
	fmt.Println(copyOfNos, nos)

}
