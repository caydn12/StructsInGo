package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	andy := Person{firstName: "Andy", lastName: "Anderson"}
	brandon := Person{"Brandon", "Branderson"}
	var cat Person
	cat.firstName = "Cat"
	cat.lastName = "Caterson"

	fmt.Println(andy)
	fmt.Println(brandon)
	fmt.Println(cat)
}
