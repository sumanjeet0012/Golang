package main

import (
	"errors"
	"fmt"
)

func main() {

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	ans2, err2 := divide2(10, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(ans2)
	}

	ans3, _ := divide(10, 2)
	fmt.Println(ans3)

}

// In general, if you need to create an error value with a simple, fixed message, use errors.New. If you need to create an error value with a dynamic message that needs to be formatted, use fmt.Errorf.

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
