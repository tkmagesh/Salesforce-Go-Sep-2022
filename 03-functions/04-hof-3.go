/* Higher order functions (functions as arguments) EXAMPLES */

package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		fmt.Println("Invocation started")
		add(100, 200)
		fmt.Println("Invocation completed")

		fmt.Println("Invocation started")
		subtract(100, 200)
		fmt.Println("Invocation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

/*
Invocation started
Add result = 300
Invocation completed

Invocation started
Subtract result = -100
Invocation completed
*/

/*
func logAdd(x, y int) {
	fmt.Println("Invocation started")
	add(x, y)
	fmt.Println("Invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("Invocation started")
	subtract(x, y)
	fmt.Println("Invocation completed")
}
*/

func logOperation(x, y int, operation func(int, int)) {
	fmt.Println("Invocation started")
	operation(x, y)
	fmt.Println("Invocation completed")
}

//3rd party library APIs
func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
