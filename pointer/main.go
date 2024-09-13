package main 

import "fmt"

func main() {
	
	var num int 
	num = 10

	var ptr *int
	ptr = &num
	fmt.Println(ptr)

	name := "sumanjeet"
	ptr2 := &name
	fmt.Println(ptr2)
	fmt.Println("The pointer points to value",*ptr2)

	var ptr3 *int
	fmt.Println(ptr3) // default value is nil

	num1 := 25
	modifyvalueusingpointer(&num1)
	fmt.Println(num1)

}

func modifyvalueusingpointer(num *int) {
	*num = *num +30
}