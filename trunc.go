package main

import (
	"errors"
	"fmt"
)

func main() {
	var inputNum float64

	inputError := errors.New("You must type in a floating point value. Please try again.")

	fmt.Println("Please type in a floating point number")

	_, err := fmt.Scanln(&inputNum)

	if err != nil {
		fmt.Println(inputError)
		var discard string
		fmt.Scanln(&discard)
		return
	}

	fmt.Println(int64(inputNum))
}
