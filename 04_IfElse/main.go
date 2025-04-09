package main

import (
	"fmt"
	"strings"
)

func main() {

	type score struct {
		name  string
		score int
	}

	scores := []score{
		{name: "phy", score: 86},
		{name: "math", score: 96},
		{name: "chem", score: 76},
	}

	fmt.Println("Student scores")

	fmt.Println("Select score to print (1-3)")
	var option string
	fmt.Scanln(&option)

	fmt.Println("student scores")
	fmt.Println(strings.Repeat("-", 14))

	var index int
	if option == "1" {
		index = 0
	} else if option == "2" {
		index = 1
	} else if option == "3" {
		index = 2
	} else {
		fmt.Println("Unknow option, defaulting to 1")
		index = 0
	}

	fmt.Println(scores[index].name, scores[index].score)

}
