/* anonymous functions */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("fn invoked")
	}()

	func(userName string) {
		fmt.Printf("Hi %s\n", userName)
	}("Magesh")

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)
}
