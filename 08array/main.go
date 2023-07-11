package main

import "fmt"

func main() {
	var fruitList [4]string
	// fruitList[0] = "Apple"
	// fruitList[1] = "Tomato"
	// fruitList[3] = "Peach"

	fmt.Println("fruit list is : ", fruitList)
	fmt.Println("fruit list length is :", len(fruitList))
	fmt.Println(fruitList[2])
	fmt.Println(fruitList[1])
}
