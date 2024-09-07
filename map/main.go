package main

import "fmt"

func main() {

	// In Go, a map (also known as a hash table or dictionary) is an unordered collection of key-value pairs.
	// It's a built-in data structure that allows you to store and retrieve data efficiently.

	studentGrade := make(map[string]int)

	studentGrade["sumanjeet"] = 100
	studentGrade["sumit"] = 90
	studentGrade["aakash"] = 80
	studentGrade["prince"] = 70
	fmt.Println(studentGrade)
	fmt.Println(studentGrade["sumanjeet"])
	fmt.Println(studentGrade["Sumanjeet"]) // case sensitive so Sumanjeet will not work, output 0.

	person := map[string]int{
		"sumanjeet": 100,
		"sumit":     90,
		"aakash":    80,
		"prince":    70,
	}
	fmt.Println(person)

	// update value
	studentGrade["prince"] = 20
	fmt.Println(studentGrade)

	// delete from map
	delete(studentGrade, "sumit")

	fmt.Println(studentGrade)

	// for loop on map
	for key, value := range studentGrade {
		fmt.Println(key, value)
	}

	// check if key exists
	if grade, ok := studentGrade["sumit"]; ok {
		fmt.Println(grade)
	} else {
		fmt.Println("Key not found")
	}
	fmt.Println(studentGrade["sumit"]) // dont know why it prints only grade and not boolean value

}
