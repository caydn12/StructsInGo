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
	fmt.Printf("%+v\n", andy)
	fmt.Println(brandon)
	fmt.Printf("%+v\n", brandon)
	fmt.Println(cat)
	fmt.Printf("%+v\n", cat)
}
