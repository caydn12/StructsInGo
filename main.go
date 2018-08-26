package main

import (
	"fmt"
)

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	james := Person{
		firstName: "James",
		lastName:  "Cameron",
		ContactInfo: ContactInfo{
			email:   "James.Cameron@gmail.com",
			zipCode: 97201,
		},
	}

	james.UpdateFirstName("Jim")

	james.Print()
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func (p *Person) UpdateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
