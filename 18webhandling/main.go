package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {

	fmt.Println("Welcome to web handling in golang")
	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of responce is %T ", responce)

	defer responce.Body.Close()

	databytes, err := ioutil.ReadAll(responce.Body)
	content := string(databytes)
	fmt.Println(content)

}
