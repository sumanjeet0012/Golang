package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	IsAdult bool
}

func main() {

	fmt.Println("Learning JSON in goLang")
	person := Person{
		Name:    "sumanjeet",
		Age:     25,
		IsAdult: true,
	}
	fmt.Println("The data of person is ", person)

	// Converting struct to JSON - JSON Marshal

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The JSON data is ", string(jsonData))

	// Converting JSON to struct - JSON Unmarshal

	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The new person is ", newPerson)

}
