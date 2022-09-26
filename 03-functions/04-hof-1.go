/* Higher order functions (functions assigned to variables) */

package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var operation func(int, int) int
	operation = func(x, y int) int {
		return x + y
	}
	fmt.Println(operation(100, 200))

	operation = func(x, y int) int {
		return x - y
	}
	fmt.Println(operation(100, 200))
}
