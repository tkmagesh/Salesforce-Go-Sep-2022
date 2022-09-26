/* Higher order functions (functions as arguments) EXAMPLES */

package main

import "fmt"

func main() {
	/*
		add(100,200)
		subtract(100,200)
	*/

	logAdd := getLoggedOperation(add)
	logAdd(100, 200)

	//logOperation(100, 200, add)

	logSubtract := getLoggedOperation(subtract)
	logSubtract(100, 200)

	//logOperation(100, 200, subtract)
}

func getLoggedOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Invocation started")
		operation(x, y)
		fmt.Println("Invocation completed")
	}
}

//3rd party library APIs
func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
