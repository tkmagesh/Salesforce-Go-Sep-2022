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
	swap( /*  */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	*x++
}

func swap( /*  */ ) {
	/*  */
}
