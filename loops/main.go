package main

import "fmt"

func main() {

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// infinite loop using for loop
	for {
		fmt.Println("Hello World")
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for range loop
	slice := []string{"a", "b", "c", "d"}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

}
