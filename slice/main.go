package main

import "fmt"

// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.
// under the hood slice is an array

func main() {

	// Create a Slice With []datatype{values}
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice2)

	var slice3 = append(slice, 11)
	fmt.Println(slice3)

	name := []string{}
	fmt.Println(name)

	// 	In Go, there are two functions that can be used to return the length and capacity of a slice:

	// len() function - returns the length of the slice (the number of elements in the slice)
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	// Create a Slice From an Array

	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice4 := arr1[0:5]
	fmt.Println(slice4)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// Create a Slice With The make() Function

	slice6 := make([]int, 5, 10)
	fmt.Println(slice6, len(slice6), cap(slice6))
	slice7 := make([]int, 5)
	fmt.Println(slice7, len(slice7), cap(slice7))
	slice7[2] = 10
	fmt.Println(slice7)

	// append() function - adds an element to the end of a slice

	// syntax - slice_name = append(slice_name, element1, element2, ...)

	slice8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice9 := append(slice8, 11, 12, 13, 14, 15)
	fmt.Println(slice9)

	// Append One Slice To Another Slice

	slice10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice11 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice12 := append(slice10, slice11...)
	fmt.Println(slice12)

}
