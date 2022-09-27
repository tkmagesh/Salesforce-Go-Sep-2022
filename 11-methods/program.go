package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func (emp Employee) WhoAmI() {
	fmt.Println("I am ", emp.Name)
}

func (emp *Employee) ChangeName(newName string) {
	emp.Name = newName
}

func main() {
	emp := Employee{Id: 100, Name: "Magesh", City: "Bangalore"}
	//(*emp).WhoAmI()
	emp.WhoAmI()
	//(&emp).ChangeName("John")
	emp.ChangeName("John")
	emp.WhoAmI()

}
