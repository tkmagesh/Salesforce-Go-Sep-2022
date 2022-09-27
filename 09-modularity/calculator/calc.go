package calculator

import "fmt"

var opCount int

func init() {
	fmt.Println("calculator initialized - 1")
}

func OpCount() int {
	return opCount
}
