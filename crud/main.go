package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

}
