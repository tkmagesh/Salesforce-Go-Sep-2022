package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	salesForce := Organization{Name: "Salesforce", City: "Newyork"}
	emp1 := Employee{Id: 100, Name: "Magesh", Org: &salesForce}
	emp2 := Employee{Id: 101, Name: "John", Org: &salesForce}

	fmt.Println(emp1)
	fmt.Println(emp2)

	emp1.Org.City = "Boston"
	fmt.Println(emp1.Org.City)
	fmt.Println(emp2.Org.City)

}
