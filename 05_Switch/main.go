package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Switch demo")

	type score struct {
		name  string
		score int
	}

	scores := []score{
		{name: "Alex", score: 86},
		{name: "Bill", score: 96},
		{name: "Dave", score: 76},
	}

	fmt.Println("Select score to print (1- 3) ")
	var option string
	fmt.Scanln(&option)

	fmt.Println("student scores")
	fmt.Println(strings.Repeat("-", 14))

	var index int

	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("Unknown option, defaulting to 1")
		index = 1
	}

	fmt.Println(scores[index].name, scores[index].score)
}
