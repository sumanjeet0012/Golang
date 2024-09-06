package main

import "fmt"

func main() {

}

func simple() {

	fmt.Println("A simple function")
}

// function with parameters and return type
func sum(x int, y int) int {
	return x + y
}

// function with named return variable
func sum2(x, y int) (result int) {
	result = x + y
	return
}

// function with multiple return types
func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

// functions with variable number of arguments

func add(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// annonymous function
func add2(a, b int) int {
	add := func(a, b int) int {
		return a + b
	}
	return add(a, b)
}

// The defer keyword is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.

func deferFunc() {
	fmt.Println("Hello")
	defer fmt.Println("World")
}
