package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loop Demo")

	i := 1
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Done!")

	// Loop with collection
	arr := [5]int{101, 102, 103, 104, 105}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Done!")
}
