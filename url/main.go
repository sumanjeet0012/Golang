package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("Learning URL in golang")
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Type of parsedURL is %T\n", parsedURL)
	fmt.Println(parsedURL)
	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)
	fmt.Println(parsedURL.Fragment)

	// Modify the URL
	parsedURL.Path = "/new/path/to/resource"
	parsedURL.RawQuery = "name=sumanjeet"

	newURL := parsedURL.String()
	fmt.Println(newURL)

}
