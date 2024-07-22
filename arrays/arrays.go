package main

import "fmt"

func main() {

	// Go Arrays
	// Arrays are used to store multple values of the same type in a single variable, instead of declaring separate
	// variables for each values.

	// Declare an array
	// 1: With the keyword:

	// Systax:
	// var array_name = [length]datatype{values}   // length is defined
	// or
	// var array_name = [...]datatype{values}      // length is inferred

	// 2: With the := sign

	// Systax:
	// array_name := [length]datatype{values}   // length is defined
	// or
	// array_name := [...]datatype{values}     // length is inferred

	// Note: The length specifies the number of elements to store in the array.
	// In Go, arrays have a fixed length. The length of the array is either defined by a number
	// or is inferred (means that the compiler decides the length of the array, based on the number of values).

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Audi"}
	fmt.Println(cars)

	// Access Elements of an array
	prices := [3]int{1, 2, 3}
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])

	// Change elements of an array
	prices[1] = 15
	fmt.Println(prices[1])

	array1 := [5]int{}              // not initialized
	array2 := [5]int{1, 2, 3}       // partially initialized
	array3 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	//It is possible to initialize only specific elements in an array.
	arr := [5]int{1: 10, 2: 40}
	fmt.Println(arr)

	//The len() function is used to find the length of an array:
	fmt.Println(len(arr))

}
