package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", "Bangalore"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bangalore",
		}
	*/

	emp := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bangalore",
	}
	fmt.Printf("%#v\n", emp)
	fmt.Println("emp name : ", emp.Name)

	empPtr := &Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bangalore",
	}
	fmt.Printf("%#v\n", empPtr)
	fmt.Println("emp (ptr) name :", (*empPtr).Name)
	fmt.Println("emp (ptr) name :", empPtr.Name)

	//empPtr2 := &Employee{}
	empPtr2 := new(Employee)
	fmt.Printf("%#v\n", empPtr2)

	emp1 := Employee{100, "Magesh", "Bangalore"}
	emp2 := Employee{100, "Magesh", "Bangalore"}
	fmt.Println(emp1 == emp2)

}
