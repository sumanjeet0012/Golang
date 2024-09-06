package main

// a directory can have a single package . so a directory can have multiple files but  all files must belong to the same package .
// for creating more packages  we can use sub directories.
// to create an executable file we use main package with main function in it.
// the main function must be in main package
// To export a function its name should start with capital letter

import (
	"fmt"
	learningPackage "hello_world2/learning"
)

func main() {
	fmt.Println("hello world ")
	learningPackage.PrintLearning()

}
