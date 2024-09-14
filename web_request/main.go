package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web requests")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while using get request")
		return
	}

	defer res.Body.Close()
	fmt.Printf("Type of responce is %T\n", res)
	// fmt.Println(res)

	// Read the responce body

	data, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		fmt.Println("error while reading responce body")
		return
	}
	fmt.Println(string(data))

	// var jsonData json.RawMessage

	// err2 := json.Unmarshal(data, &jsonData)
	// if err2 != nil {
	// 	fmt.Println("error while unmarshalling")
	// 	return
	// }
	// jsonData1 := string(jsonData)
	// fmt.Println(jsonData1)
}
