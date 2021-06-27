/*
Write a program which
- x prompts the user to first enter a name,
- and then enter an address.
- Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
- Your program should use Marshal() to create a JSON object from the map,
- and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var NAME_PROMPT = "Please enter your name and press ENTER."
var ADDRESS_PROMPT = "Please enter your address and press ENTER."
var INPUT_ERROR = "Valid JSON with your data was not created. Please try again."

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	MakeJson(scanner)
}

func MakeJson(scanner *bufio.Scanner) {
	userMap := map[string]string{}

	fmt.Println(NAME_PROMPT)

	scanner.Scan()

	name := scanner.Text()

	userMap["name"] = name

	fmt.Println(ADDRESS_PROMPT)

	scanner.Scan()

	address := scanner.Text()

	userMap["address"] = address

	userJson, err := json.Marshal(userMap)

	if err != nil {
		fmt.Fprintln(os.Stderr, INPUT_ERROR, err)
	}

	fmt.Println(string(userJson))
}
