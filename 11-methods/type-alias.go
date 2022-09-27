package main

import "fmt"

type MyStr string

func (str MyStr) Length() int {
	return len(str)
}

func main() {
	s := MyStr("This is a string")
	fmt.Println(s.Length())
}
