package main

import (
	"fmt"
	//"modularity-demo/calculator"
	// _ "modularity-demo/calculator" // importing ONLY to run the init functions
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("app executed!")
	color.Green("app started!")
	fmt.Println(calc.Add(100, 200))
	run()
	fmt.Println("OpCount = ", calc.OpCount())
	fmt.Println(utils.IsEven(21))
	color.Red("app completed!")
}
