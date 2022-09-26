/* Higher order functions (functions as return values) */

package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
