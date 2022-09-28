package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type ShapeWithArea interface {
	Area() float32
}

func main() {
	//var x interface{}
	var x any
	x = "This is a string"
	//x = 10.456
	//x = true
	//x = struct{}{}
	fmt.Println(x)

	//x = 100
	//y := x + 100
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}
	x = Circle{12}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 = ", val+200)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	case struct{}:
		fmt.Println("x is a struct{}")
	case Circle:
		fmt.Println("x is a circle with radius :", val.Radius)
	case ShapeWithArea:
		fmt.Println("Area = ", val.Area())
	default:
		fmt.Println("unknown type")
	}

}
