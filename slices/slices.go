package main

import "fmt"

// Go Slices

// Slices are similar to arrays, but more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.

// However, unlike arrays, the length of a slice can grow and shrik as you see fit.

// In Go, there are several ways to create a slice:
// Using the []datatype{values} format.
// Create a slice from an array.
// Using the make() function.

func main() {

	// Create a slice with []datatype{values}
	// Syntax:
	// slice_name :=[]datatype{values}

	myslices := []int{}

	// Before add value
	fmt.Println(myslices)      // Empty slice
	fmt.Println(cap(myslices)) // slice capacity
	fmt.Println(len(myslices)) // slice length

	// Append value
	myslices = append(myslices, 10)
	fmt.Println(myslices)
	fmt.Println(cap(myslices)) // slice capacity
	fmt.Println(len(myslices)) // slice length

	myslice1 := []int{1, 2, 3}
	fmt.Println(myslice1)      // slice
	fmt.Println(cap(myslice1)) // slice capacity
	fmt.Println(len(myslice1)) // slice length

	// Create a slice from an array
	// Syntax:
	// var myarray = [length]datatype{values}	// An Array
	// myslice := myarray[start:end]			// A slice made from th array

	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	mySlice := arr1[2:4]
	fmt.Println(mySlice)

	// Create a slice with the make() function
	// Syntax:
	// slice_name := make([]type, length, capacity)  // If the capacity parameter is not defined, it will be equal to length

	slice3 := make([]int, 0)
	slice3 = append(slice3, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice3)

	// Create slices

}
