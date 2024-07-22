package main

import "fmt"

// If a variable should have a fixed value that cannot be changed,
// you can use the const keyword.

// The const keyword declares the variable as "constant",
// which means that it is unchangeable and read-only.

// Syntax
//  const CONSTNAME type = value

const SERVER_URL = "http://localhost"

func main() {
	fmt.Println(SERVER_URL)

	// 	Constant Types
	// There are two types of constants:

	// Typed constants
	// Untyped constants

	// 	Typed Constants
	// Typed constants are declared with a defined type:

	const A int = 1
	fmt.Println(A)

	// 	Untyped Constants
	// Untyped constants are declared without a type:

	const B = 20
	fmt.Println(B)

}
