package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("please rate")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)    // here input value is "4"
	fmt.Println("Thanks for rating, ", "44\n\n") // here input value is "4"

	numRating, err := strconv.ParseFloat(input, 64) // but here input value is "4\n", why?

	if err != nil {
		fmt.Println(err)
		fmt.Println("test")
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
