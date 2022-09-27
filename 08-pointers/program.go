package main

import "fmt"

func main() {

	var no int = 10
	//noPtr := &no //(address of no OR pointer to no)
	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)

	//deferencing (accessing the value using the pointer)
	fmt.Println(*noPtr)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)

	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, nos = ", nos)
	sort(&nos)
	fmt.Println("After sorting, nos = ", nos)
}

func increment(x *int) {
	*x++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func sort(nos *[5]int) {
	for i := 0; i < len(nos)-1; i++ {
		for j := i + 1; j < len(nos); j++ {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
