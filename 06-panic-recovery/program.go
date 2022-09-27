/* panic & recovery */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something went wrong. ")
			fmt.Println(err)
		}
		fmt.Println("Thank you")
	}()
	quotient, remainder, err := divideClient(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Please do not divide by zero")
	} else {
		fmt.Println(quotient, remainder)
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)

	}
	quotient, remainder = x/y, x%y
	return
}
