package main

import "fmt"

func main() {

	// The defer keyword can only be used with a function call, not with an assignment or part of an expression.
	// deffered statements always execute in LIFO order
	// deffered staatements always execute regardless of whether an error occurs or not and reguardless of the exitpoint in the code

	fmt.Println("Starting of the code ")
	defer fmt.Println("Middle of the code ")
	fmt.Println("End of the code ")

	a := 1
	if a == 10 {
		defer fmt.Println("10")
	}

	sum := add(10, 20)
	fmt.Println(sum)

}

func add(a, b int) int {

	return a + b

}
