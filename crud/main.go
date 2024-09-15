package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	fmt.Println("Learning CRUD in golang")

	// GET request

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while using get request", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error while using get request", res.StatusCode)
		return
	}

	// Read the responce body

	// data, err1 := ioutil.ReadAll(res.Body)
	// if err1 != nil {
	// 	fmt.Println("error while reading responce body", err1)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error while unmarshalling", err)
		return
	}
	fmt.Println(todo)

	// POST request

	todo1 := Todo{
		UserId:    1,
		Title:     "Buy milk",
		Completed: false,
	}

	// Convert struct to JSON

	jsonData1, err2 := json.Marshal(todo1)
	if err != nil {
		fmt.Println("error while marshalling", err2)
		return
	}
	fmt.Println(jsonData1)
	fmt.Println(string(jsonData1))

	// Convert JSON to string

	jsonString := string(jsonData1)
	fmt.Println(jsonString)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	fmt.Println(jsonReader)
	fmt.Println(*jsonReader)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request

	res3, err3 := http.Post(myUrl, "application/json", jsonReader)

	if err3 != nil {
		fmt.Println("error while using post request", err3)
		return
	}

	defer res3.Body.Close()

	data, _ := ioutil.ReadAll(res3.Body)

	fmt.Println("Status is", res3.Status)
	fmt.Println("Response is", string(data))

}
