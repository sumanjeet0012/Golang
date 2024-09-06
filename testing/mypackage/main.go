package main

import (
	"fmt"
	"testingpackage/anotherpackage"
)

func main() {

	fmt.Println("hello world from main package ")

	anotherpackage.PrintAnother()
}
