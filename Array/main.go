package main

import "fmt"

func main() {

	//Array
	var arr [3]int

	arr = [3]int{1, 2, 3}

	fmt.Println(arr[1])
	arr[1] = 86
	fmt.Println(arr[1])

	fmt.Println(arr)

	//length of array
	fmt.Println(len(arr))

	//string array
	strArr := [4]string{"alex", "cedric", "bill", "Dave"}

	arr2 := strArr
	fmt.Println(arr2)

	strArr[1] = "foo"
	fmt.Println(strArr)
	fmt.Println(arr2)

	//------------

	// Slices
	// Todo: add code here

	//------------
	// Maps
	// Todo: add code here
	var fruit map[string]int
	fmt.Println(fruit)
	fruit = map[string]int{"Mango": 12, "Apple": 6}
	fmt.Println(fruit)

	fruit["banana"] = 24
	fmt.Println(fruit)

	fmt.Println(fruit["banana"])

	delete(fruit, "Apple")
	fmt.Println(fruit)

}
