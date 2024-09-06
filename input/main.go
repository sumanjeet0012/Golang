package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var name string
	var age int
	fmt.Scan(&name, &age)
	fmt.Println(name, age)

	var fullName string

	fmt.Println("Enter your full name")
	fmt.Scanln(&fullName)

	fmt.Println("Hello", fullName)

	// for more complex input we can use bufio package
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	bufname, _ := reader.ReadString('\n')
	fmt.Println("Hello", bufname)
}
