/* variadic functions */

package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40))
	//fmt.Println(sum())
}

func sum(no int, nos ...int) int {
	result := no
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
