package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one one two two two three three"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str2 := "        hello   world    "
	trim := strings.TrimSpace(str2)
	fmt.Println(str2)
	fmt.Println(trim)

	name := "sumanjeet"
	upper := strings.ToUpper(name)
	fmt.Println(upper)

	firstName := "Sumanjeet"
	lastName := "Singh"
	join := strings.Join([]string{firstName, "Kumar", lastName}, " ")
	fmt.Println(join)

}
