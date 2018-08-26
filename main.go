package main

import (
	"fmt"
)

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func main() {
	james := Person{
		firstName: "James",
		lastName:  "Cameron",
		contactInfo: ContactInfo{
			email:   "James.Cameron@gmail.com",
			zipCode: 97201,
		},
	}

	fmt.Printf("%+v", james)
}
