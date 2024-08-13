package main

import "fmt"

func main() {
	// Slices are like arrays that you can dynamically resize and pass to functions more efficiently.
	// Maps are associative arrays, unordered list of key/value pairs that allow you to efficiently and
	//quickly look up values for a unique key. (like dictionaries in python)

	var s = make([]string, 0)
	var m = make(map[string]string)

	s = append(s, "some string")
	m["some key"] = "some value"

	fmt.Println(s)
	fmt.Println(m)
}
