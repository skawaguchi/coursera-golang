package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Find(scanner)
}

var VALID_MSG = "Found!"
var INVALID_MSG = "Not Found!"

func Find(scanner *bufio.Scanner) {

	fmt.Println("Please enter a string starting with 'i', ending in 'n', and containing 'a'.")

	scanner.Scan()

	chars := strings.Split(scanner.Text(), "")

	isValid := hasValidString(chars)

	if isValid {
		fmt.Println(VALID_MSG)
	} else {
		fmt.Println(INVALID_MSG)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading stdin", err)
	}
}

var EXPECTED_FIRST_CHAR = "i"
var EXPECTED_MID_CHAR = "a"
var EXPECTED_LAST_CHAR = "n"

var STARTING_ERROR_MSG = "you must start your string with 'i'"
var MID_ERROR_MSG = "your string must have an 'a' somewhere in the middle of it"
var LAST_ERROR_MSG = "your string must have end with 'n'"

func hasValidString(chars []string) bool {
	startingError := errors.New(STARTING_ERROR_MSG)
	midError := errors.New(MID_ERROR_MSG)
	endError := errors.New(LAST_ERROR_MSG)

	hasCorrectMidChar := false

	for i := 0; i < len(chars); i++ {
		char := chars[i]

		if i == 0 {
			if char != EXPECTED_FIRST_CHAR {
				fmt.Println(startingError)
				return false
			}
		} else if i == len(chars)-1 {
			if !hasCorrectMidChar {
				fmt.Println(midError)
				return false
			} else if char != EXPECTED_LAST_CHAR {
				fmt.Println(endError)
				return false
			}
		} else if i < len(chars)-1 {
			if char == EXPECTED_MID_CHAR {
				hasCorrectMidChar = true
			}
		}
	}
	return true
}
