package main

import "fmt"

func main() {
	x := int(10)

	// You dont wrap the conditional check in parentheses
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	var y string = "foo"
	// For conditionals involving more than two choices, Go provides a switch statement
	// Unlike many other languages, your cases dont have to include a break statement
	// Go will not execute more than one matching or default case
	switch y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	fmt.Println("\nLoop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nLoop through a collection Slice or Map")
	nums := []int{2, 4, 6, 8}
	// range returns two values: the current index and a copy of the current value at that index.
	// The index can be replace by a _ if you don't need it
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}
