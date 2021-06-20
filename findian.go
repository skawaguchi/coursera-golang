/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program,
“findian.go”.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var EXPECTED_FIRST_CHAR = "i"
var EXPECTED_MID_CHAR = "a"
var EXPECTED_LAST_CHAR = "n"

func main() {
	startingError := errors.New("You must start your string with 'i'.")
	midError := errors.New("Your string must have an 'a' somewhere in the middle of it.")
	endError := errors.New("Your string must have end with 'n'.")

	fmt.Println("Please enter a string starting with 'i', ending in 'n', and containing 'a'.")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	chars := strings.Split(scanner.Text(), "")

	hasCorrectFirstChar := false
	hasCorrectMidChar := false
	hasCorrectLastChar := false

	for i := 0; i < len(chars); i++ {
		char := chars[i]

		if i == 0 {
			if char != EXPECTED_FIRST_CHAR {
				fmt.Println(startingError)
				return
			} else {
				hasCorrectFirstChar = true
			}
		} else if i == len(chars)-1 {
			if hasCorrectMidChar != true {
				fmt.Println(midError)
				return
			} else if char != EXPECTED_LAST_CHAR {
				fmt.Println(endError)
				return
			} else {
				hasCorrectLastChar = true
			}
		} else if i < len(chars)-1 {
			if char == EXPECTED_MID_CHAR {
				hasCorrectMidChar = true
			}
		}
	}

	if hasCorrectFirstChar && hasCorrectMidChar && hasCorrectLastChar {
		fmt.Println("You've entered a valid string. Yay!")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading stdin", err)
	}
}
