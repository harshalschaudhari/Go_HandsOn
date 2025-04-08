package main

import (
	"fmt"
	"strings"
)

func main() {

	// declare variable
	var myName string
	myName = "Cedric"
	fmt.Println(myName)

	//declare and initialize
	var myFirstName string = "Alex"
	fmt.Println(myFirstName)

	//Short hand variable declare and assignment
	name, score := "Harshal, Chaudhari", 86

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score)
}
