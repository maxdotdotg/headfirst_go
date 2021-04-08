package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

func main() {
	// create an instance of Subscriber,
	// then supply attributes
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	// these only work when the HomeAddress field in the
	// subsciber struct is defined

	//s.HomeAddress.Street = "234 5th street"
	//s.Hfmt.Println(s.HomeAddress)
	//s.Hfmt.Printf("%#v\n", s.HomeAddress)
	//s.Hfmt.Printf("%#v\n", s)

	// create an instance of Subscriber as a struct literal
	newSub := magazine.Subscriber{Name: "kelly", Rate: 3.99}
	fmt.Println(newSub.Name)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	var testAddress magazine.Address
	testAddress.Street = "123 4th street"
	testAddress.City = "new jack city"
	testAddress.State = "OHNO"
	testAddress.PostalCode = "ABC easy as 123"
	fmt.Println(testAddress)

	// accessing nested struct attributes with anonymous fields
	newEmployee := magazine.Employee{Name: "not kelly"}
	newEmployee.Street = "some street somewhere"
	fmt.Println(newEmployee.Address)
	fmt.Printf("%#v\n", newEmployee)

}
