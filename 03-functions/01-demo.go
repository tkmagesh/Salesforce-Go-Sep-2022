/* function basics */
package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Println(getGreetMsg("Magesh"))
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d \n", quotient)
}

/* 01. basic function */
func fn() {
	fmt.Println("fn invoked")
}

/* 02. function with argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* 03. function with argument and return value */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!", userName)
}

/* 04. function with multiple arguments and 1 return value */
/*
func add(x int, y int) int {
	return x + y
}
*/
func add(x, y int) int {
	return x + y
}

/* 05. function with multiple arguments and mutiple return values */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/* 06. function with named return values */
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

/*
func divide(x, y int) (quotient, remainder int) {
	remainder = x % y
	quotient = x / y
	return
}
*/
func divide(x, y int) (quotient, remainder int) {
	remainder, quotient = x%y, x/y
	return
}
