/*
Write a program which:
- x prompts the user to enter integers
- stores the integers in a sorted slice.
- The program should be written as a loop.
- Before entering the loop, the program should create an empty integer slice of size (length) 3.
- During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
- The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
- The slice must grow in size to accommodate any number of integers which the user decides to enter.
- The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var PROMPT = "Please enter an integer and type ENTER. To quit, enter 'x' and type ENTER."
var EXIT_MSG = "Exiting the program. Bye!"
var INPUT_ERROR = "You did not type an integer as a valid entry. Please try again."

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Slice(scanner)
}

func Slice(scanner *bufio.Scanner) {
	slice := make([]int, 3)
	slice = slice[:0]

	fmt.Println(PROMPT)

	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "x" {
			fmt.Println(EXIT_MSG)
			break
		}

		inputInt, strConvErr := strconv.Atoi(input)

		if strConvErr != nil {
			fmt.Println(INPUT_ERROR)
		} else {
			slice = append(slice, inputInt)
			sort.Ints(slice)
			fmt.Println(slice)
		}
		fmt.Println(PROMPT)
	}

}
