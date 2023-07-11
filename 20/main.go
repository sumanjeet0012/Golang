package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=abcd1234"

func main() {

	fmt.Println("Welcome to web handling in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)

}
