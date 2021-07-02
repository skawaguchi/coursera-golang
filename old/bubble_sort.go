/*
- The program should prompt the user to type in a sequence of up to 10 integers.
- The program should print the integers out on one line, in sorted order, from least to greatest.
- Use your favorite search tool to find a description of how the bubble sort algorithm works.
- As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.
- A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation.
- Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
- The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
- Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var PROMPT = "Please enter up to 10 integers, separated by spaces, and type ENTER to sort them."
var PROMPT_LIMIT_REACHED = "You have entered the limit of %d numbers. Taking the first %d.\n"
var INPUT_ERROR = "You did not type an integer as a valid entry. Please try again."
var FINAL_OUTPUT_MSG = "Sorted output:"
var INPUT_LIMIT = 10

func main() {
	fmt.Println(PROMPT)

	scanner := bufio.NewScanner(os.Stdin)

	input := getInput(scanner)

	sorted := BubbleSort(input)

	fmt.Println(FINAL_OUTPUT_MSG, sorted)
}

func getInput(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Fields(scanner.Text())

	inputtedInts := []int{}

	if len(input) > INPUT_LIMIT {
		input = input[:INPUT_LIMIT]
		fmt.Printf(PROMPT_LIMIT_REACHED, INPUT_LIMIT, INPUT_LIMIT)
	}

	for _, char := range input {
		inputInt, strConvErr := strconv.Atoi(char)

		if strConvErr != nil {
			fmt.Println(INPUT_ERROR)
		} else {
			inputtedInts = append(inputtedInts, inputInt)
		}
	}

	return inputtedInts
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr = Swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func Swap(arr []int, first int, second int) []int {
	arr[first], arr[second] = arr[second], arr[first]
	return arr
}
