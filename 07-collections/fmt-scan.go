package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter the name :")
		fmt.Scanln(&name)
		fmt.Println(name)
	*/
	/*
		var no int
		fmt.Println("Enter the no :")
		n, err := fmt.Scan(&no)
		fmt.Println(n, err)
		fmt.Println(no)
	*/

	var n1, n2 int
	if n, err := fmt.Scanln(&n1, &n2); n < 2 {
		fmt.Println("input error : ", err)
	} else {
		fmt.Println(n1 + n2)
	}

}
