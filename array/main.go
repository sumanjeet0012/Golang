package main

import "fmt"

func main() {
	fmt.Println("we are learning Array in Golang")

	// var name[5]string

	// //  0 1 2 3 4
	// name[0] = "prince"
	// name[2] = "Aakash"

	// fmt.Println("Names of Person is :", name)

	// var numbers = [8]int{1, 2, 3, 4, 5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of Numbers array is :", len(numbers))

	// fmt.Println("value of name at 2 index is :", name[2])

	// ""
	var name [5]string
	name[2] = "Prince"
	name[0] = "aaksh"

	fmt.Println("name is :", name)
	fmt.Printf("name Array is  %q\n", name)

	// array is initialized with default value
	var price [5]int
	fmt.Println("Price is :", price)

	// easy way to remember array
	var namess [5]string = [5]string{"aakash", "prince", "khan"}
	fmt.Println("Name is :", namess)
}
