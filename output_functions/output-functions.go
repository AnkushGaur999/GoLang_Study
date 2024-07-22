package main

import "fmt"

// Go has three functions to output text:

// Print()
// Println()
// Printf()

func main() {

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	var a, b string = "Hello", "World"

	fmt.Print(a, "\n")
	fmt.Print(b, "\n")

	// \n creates new lines.

	// It is also possible to only use one Print() for printing multiple variables.
	fmt.Print(a, "\n", b)

	// If we want to add a space between string arguments, we need to use " ":
	fmt.Print(a, " ", b)

	// 	The Println() Function: ->
	// The Println() function is similar to Print() with the difference that a
	// whitespace is added between the arguments, and a newline is added at the end:
	fmt.Print(a, b)

	// The Printf() function

	// 	The Printf() function first formats its argument based on the given formatting verb and then prints them.

	// Here we will use two formatting verbs:

	// %v is used to print the value of the arguments.
	// %T is used to print the type of the arguments.

	fmt.Printf("a has value: %v and type %T\n", a, a)
	fmt.Printf("b has value: %v and type %T\n", b, b)


// 	 Verb			Description
//   %v			Prints the value in the default format
//   %#v		Prints the value in Go-syntax format
//   %T			Prints the type of the value
//   %%			Prints the % sign

}
