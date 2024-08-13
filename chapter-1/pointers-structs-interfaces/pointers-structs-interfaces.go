package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Dog struct{}

type Friend interface {
	SayHello()
}

func (d *Dog) SayHello() {
	fmt.Println("Woof, Woof")
}

func (p *Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	// Pointer
	fmt.Println("Pointer")
	var count = int(42)
	// & - gives us the address in memory of a variable
	ptr := &count
	fmt.Println("Pointer value: ", *ptr)
	//  * - dereference the address
	*ptr = 100
	fmt.Println("Pointer new value: ", count)

	// Struct
	// Types and fields that begin with a capital letter are exported and acessible outside the package
	// Whereas those starting with a lowercase are private, accessible only within the package
	fmt.Println("\nStruct")
	var guy = new(Person)
	guy.Name = "Joe"
	guy.Age = 42
	guy.SayHello()

	// Interface
	// Interface type are like a blueprint or contract. This blueprint defines an expected set of actions that any
	// concrete implementation must fulfill in order to be considered a type of that interface
	fmt.Println("\nInterface")
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}
