package main

import "fmt"

func main() {
	//nos := []int{}
	/*
		nos := []int{10}
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 20)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 30)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 40)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 50)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos) */

	//initialization
	/*
		nos := make([]int, 3)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 10)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/

	//ONLY allocation
	/*
		nos := make([]int, 0, 10)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		//nos[0] = 10
		nos = append(nos, 10)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/

	//initialization + allocation
	nos := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
}
