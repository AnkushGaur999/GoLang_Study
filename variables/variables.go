package main

import "fmt"

func main() {

	// Go variable type

	// int - store integers (whole numbers) suhc as 123 or -123
	// float - store floating point numbers, with decimals, such as 19.99 0r -19.99
	// string - store text, such as "Hello World". String values are surrounded by double quotes
	// bool - store values

	// Declaring (Cretaing) variables

	// 1. with the var keyword
	// Syntax
	// var varibalename type = value

	var name string = "Ankush"
	fmt.Println(name)

	// Note: You always have to specify either type or value (or both).

	// 2. With the := sign:
	// Use the := sign, followed by the variable value:
	// Syntax
	//  varibalename  := value

	age := 20
	fmt.Println(age)
	// Variable Declaration Without Initial Value
	// In Go, all variables are initialized. So,
	// if you declare a variable without an initial value,
	//  its value will be set to the default value of its type:
	var a string
	var b int
	var c bool

	fmt.Println(a) // default value for string = ""
	fmt.Println(b) // default value for int = 0
	fmt.Println(c) // default value for bool = false

	// 	Value Assignment After Declaration
	// It is possible to assign a value to a variable after it is declared.
	// This is helpful for cases the value is not initially known.
	var student1 string
	student1 = "John"
	fmt.Println(student1)

	// Difference Between var and :=
	// ----------------( var )----------------			---------------( := )----------------
	// Can be used inside and outside of functions.		Can only be used inside functions.
	// Variable declaration and value assignment    	Variable declaration and value assignment
	// can be done separately.							cannot be done separately(must be done in the same line).

	// In Go, it is possible to declare multiple variables in the same line.

	var w, x, y, z int = 1, 3, 5, 7

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//If the type keyword is not specified, you can declare different types of variables in the same line:
	var l, m = 6, "Hello"
	n, o := 7, "World!"

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	// Multiple varibale declaration can also be grouped together into a block for greater readability.

	var(
		pName string
		pAge = 22
		pDesc = "Description"
	)

	fmt.Println(pName)
	fmt.Println(pAge)
	fmt.Println(pDesc)


}
