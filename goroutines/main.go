package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Learning goroutines in golang")
	go hello()
	go hi()
	time.Sleep(2 * time.Second)

}

func hello() {

	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}

}

func hi() {

	for i := 0; i < 5; i++ {
		fmt.Println("hi")
	}

}
