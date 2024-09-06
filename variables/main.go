package variables

import "fmt"

func main() {
	fmt.Println("variables")

	// writing data type is optional in go

	var name string = "sumanjeet"
	fmt.Println(name)

	var age = 25
	fmt.Println(age)

	var isCool = true
	fmt.Println(isCool)

	var size float32 = 1.3
	fmt.Println(size)

	var dedicated bool = true
	fmt.Println(dedicated)

	const pi = 3.14
	fmt.Println(pi)

	// short hand declaration
	personName := "sumanjeet"
	fmt.Println(personName)

	// public and private variable
	var PublicName string = "sumanjeet"
	fmt.Println(PublicName)

	var privateName string = "suraj"
	fmt.Println(privateName)

}
