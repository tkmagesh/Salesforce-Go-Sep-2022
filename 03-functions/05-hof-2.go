/* Higher order functions (functions as arguments) EXAMPLES */

package main

import (
	"fmt"
	"time"
)

type Operation func(int, int)

func main() {
	/*
		add(100,200)
		subtract(100,200)
	*/

	/*
		logAdd := getLoggedOperation(add)
		logAdd(100, 200)
	*/

	//logOperation(100, 200, add)

	/*
		logSubtract := getLoggedOperation(subtract)
		logSubtract(100, 200)
	*/

	//logOperation(100, 200, subtract)

	/* profiledAdd := getProfiledOperation(add)
	profiledAdd(100, 200)

	profiledSubtract := getProfiledOperation(subtract)
	profiledSubtract(100, 200) */

	add := getProfiledOperation(getLoggedOperation(add))
	subtract := getProfiledOperation(getLoggedOperation(subtract))

	add(100, 200)
	subtract(100, 200)
}

//func getLoggedOperation(operation func(int, int)) func(int, int) {
func getLoggedOperation(operation Operation) Operation {
	return func(x, y int) {
		fmt.Println("Invocation started")
		operation(x, y)
		fmt.Println("Invocation completed")
	}
}

func getProfiledOperation(operation Operation) Operation {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Println("Elapsed =", elapsed)
	}
}

//3rd party library APIs
func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
