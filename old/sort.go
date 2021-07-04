/*
- The goal of this activity is to explore the use of threads by creating a program for sorting integers
- uses four goroutines to create four sub-arrays
- then merge the arrays into a single array.

- Students will receive five points if the program sorts the integers
- five additional points if there are four goroutines that each prints out a set of array elements that it is storing.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var ARRAY_LENGTH = 12
var TARGET_SUB_ARRAYS = 4.0
var PROMPT = "Please enter up to %d integers separated by SPACES, and hit ENTER to sort.\n"
var INPUT_CHAR_ERROR_MSG = "You entered a non-integer. %s Please try again.\n"

func main() {
	fmt.Printf(PROMPT, ARRAY_LENGTH)

	scanner := bufio.NewScanner(os.Stdin)

	inputInts := getInput(scanner)

	sortedInts := sortSlice(inputInts)
	fmt.Println("Sorted:", sortedInts)
}

func getInput(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Fields(scanner.Text())

	inputtedInts := []int{}

	for _, char := range input {
		inputInt, strConvErr := strconv.Atoi(char)

		if strConvErr != nil {
			fmt.Printf(INPUT_CHAR_ERROR_MSG, inputtedInts)
		} else {
			inputtedInts = append(inputtedInts, inputInt)
		}
	}

	return inputtedInts
}

func sortSlice(slice []int) []int {
	var wg sync.WaitGroup
	var channel = make(chan []int)
	subSliceSize := int(math.Ceil(float64(len(slice)) / TARGET_SUB_ARRAYS))
	rtnSlice := []int{}
	id := 1
	for startIndex := 0; startIndex < len(slice); startIndex += subSliceSize {
		endIndex := startIndex + subSliceSize
		if endIndex > len(slice) {
			endIndex = len(slice)
		}

		sortSubArray(channel, id, &wg, slice[startIndex:endIndex])
		subArray := <-channel
		rtnSlice = append(rtnSlice, subArray...)
		id += 1
	}
	wg.Wait()

	sort.Ints(rtnSlice)

	return rtnSlice
}

func sortSubArray(channel chan<- []int, id int, wg *sync.WaitGroup, slice []int) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Worker ", id, " sorting: ", slice)
		sort.Ints(slice)
		channel <- slice
	}()
}
