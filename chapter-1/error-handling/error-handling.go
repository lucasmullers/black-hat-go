package main

import (
	"errors"
	"fmt"
)

// type MyError string

// func (e MyError) Error() string {
// 	return string(e)
// }

func foo(x int) error {
	if x == 1 {
		return nil
	} else {
		return errors.New("Some Error Occured")
	}
}

func main() {
	// Go does not include a syntax for try / catch / finally error handling
	// Go defines a built-in error type with the following interface declaration: type error interface { Error() string }
	fmt.Println("Error Handling")

	err := foo(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")

}
