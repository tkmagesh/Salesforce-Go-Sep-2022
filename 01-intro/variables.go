package main

import "fmt"

var no int

// var no int = 100
// var no = 100
// no := 100 //=> NOT allowed at the package level

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		//using the compiler's type inference
		var msg = "Hello World!"
	*/

	msg := "Hello World!" //=> ":=" can be used only for a function scoped variable (NOT valid at package scope)
	fmt.Println(msg)

	// declaring multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x = 100
		var y = 200
		var str = "Sum of %d and %d is %d\n"
		var result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

	/* ======================================== */
	//constants
	// const PI float32 = 3.14
	// const PI = 3.14

	/*
		const (
			PI float32 = 3.14
			NO int     = 100
		)
	*/

	/*
		const (
			PI = 3.14
			NO = 100
		)
	*/
	const (
		PI, NO = 3.14, 100
	)

	/* iota */

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	const (
		red = iota * 2
		green
		_
		blue
	)

	fmt.Printf("red= %d, green = %d, blue = %d\n", red, green, blue)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)

	//complex type
	var n1 complex64 = 4 + 2i
	fmt.Println(n1)
	fmt.Printf("real = %v, imaginary = %v\n", real(n1), imag(n1))

	var n2 complex64 = 2 + 5i
	fmt.Println(n1 + n2)
}
