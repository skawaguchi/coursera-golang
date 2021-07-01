/*
- Write a program which reads information from a file
- represents it in a slice of structs.
- Assume that there is a text file which contains a series of names.
	- Each line of the text file has a first name and a last name, in that order
	- separated by a single space on the line.
- Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
- Each field will be a string of size 20 (characters).

- Your program should prompt the user for the name of the text file.
- Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
- Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
- After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

- Submit your source code for the program, “read.go”.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var FILE_OPEN_ERROR = "There was a problem OPENING the file."
var FILE_READ_ERROR = "There was a problem READING the file."
var INPUT_PROMPT = "Please type in the relative path to the name file."
var TRUNC_WARNING = "The %s name %q was longer then the limit of %d characters. It has been truncated to %q."
var CHAR_LIMIT = 20

type Name struct {
	fname string
	lname string
}

func main() {
	nameSlice := make([]Name, 0)
	inputScanner := bufio.NewScanner(os.Stdin)
	filePath := promptFilePathInput(inputScanner)
	populatedNameSlice := getFileContent(filePath, nameSlice)
	printNames(populatedNameSlice)
}

func printNames(nameSlice []Name) {
	for _, name := range nameSlice {
		fmt.Println(name.fname, name.lname)
	}
}

func promptFilePathInput(inputScanner *bufio.Scanner) string {
	fmt.Println(INPUT_PROMPT)

	inputScanner.Scan()

	filePath := inputScanner.Text()

	return filePath
}

func getFileContent(filePath string, nameSlice []Name) []Name {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(FILE_OPEN_ERROR, err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		names := strings.Split(fileScanner.Text(), " ")
		fname, lname := getValidNames(names[0], names[1])
		nameStruct := getName(fname, lname)
		nameSlice = append(nameSlice, nameStruct)
	}

	scannerErr := fileScanner.Err()
	if scannerErr != nil {
		log.Fatal(FILE_READ_ERROR, scannerErr)
	}

	return nameSlice
}

func getValidNames(fname string, lname string) (validFname string, validLname string) {
	truncatedFName := fname
	truncatedLName := lname
	if len(fname) > CHAR_LIMIT {
		truncatedFName = fname[:CHAR_LIMIT]
		log.Printf(TRUNC_WARNING, "first", fname, CHAR_LIMIT, truncatedFName)
	}
	if len(lname) > CHAR_LIMIT {
		truncatedLName = lname[:CHAR_LIMIT]
		log.Printf(TRUNC_WARNING, "last", lname, CHAR_LIMIT, truncatedLName)
	}

	return truncatedFName, truncatedLName
}

func getName(fname string, lname string) Name {
	newName := Name{fname: fname, lname: lname}
	return newName
}
