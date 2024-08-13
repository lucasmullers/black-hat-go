package main

import "fmt"

func main() {
	// Primitive Data Types:
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64
	// byte
	// float32, float64
	// complex64, complex128
	// bool
	// string
	// rune
	// uintptr

	// If you don't specify a type, Go will infer it from the value you give it.
	// Both ways of declaring a variable are valid.
	var x = "Hello, World!"
	z := int(42)

	fmt.Println(x, z)
}
