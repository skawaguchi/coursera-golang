package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ENTER_COMMAND_MSG = "Type in a command and hit ENTER. Possible commands: %s. (Type `x` to EXIT.)\n"
var ENTER_NAME_MSG = "Type in a name for your animal. (Type `x` to EXIT.)"
var ENTER_ANIMAL_TYPE_MSG = "Type in the type of animal you want to create and hit ENTER.\nThe choices for types of animals are: %s. (Type `x` to EXIT.)\n"
var ENTER_QUERY_NAME_MSG = "Type in the name for the animal you have previously created. (Type `x` to EXIT.)"
var ENTER_ANIMAL_COMMAND_MSG = "Type in the command for your animal. The choices are: %s. (Type `x` to EXIT.)\n"

var INVALID_INPUT_MSG = "You must enter %d input command(s). Try again.\n"
var INVALID_COMMAND_MSG = "You typed in %q as the command. This is invalid. Try again using a valid command. Your options are %s.\n"
var INVALID_ANIMAL_NAME_MSG = "You must enter a name of at least %d characters. Try again.\n"
var INVALID_ANIMAL_TYPE_MSG = "You typed in %q as the animal type. This is invalid. Please try again. Your options are %s.\n"
var INVALID_ANIMAL_ACTIVITY_MSG = "You typed in %q as the animal activity. This is invalid. Please try again. Your options are %s.\n"
var INVALID_ACTIVITY_MSG = "You typed in %q as the animal activity. This is invalid. Your options are %s.\n"
var INVALID_ANIMAL_MSG = "You typed in %q as the animal, but that animal doesn't exist.\n"

var CREATED_ANIMAL_MSG = "Created it!"
var EXIT_MSG = "Bye Bye!"
var MIN_ANIMAL_NAME_LENGTH = 1

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

const (
	INVALID_COMMAND = "invalid"
	NEW_ANIMAL      = "newanimal"
	QUERY           = "query"
)

const (
	EAT   = "eat"
	MOVE  = "move"
	SPEAK = "speak"
)

var commandList = []string{NEW_ANIMAL, QUERY}
var animalTypeList = []string{COW, BIRD, SNAKE}
var animalActivityList = []string{EAT, MOVE, SPEAK}
var commandPrintList = getListToPrint(commandList)
var animalTypePrintList = getListToPrint(animalTypeList)
var animalActivityPrintList = getListToPrint(animalActivityList)
var userAnimals = map[string]Animal{}

const (
	BIRD  = "bird"
	FLY   = "fly"
	WORMS = "worms"
	PEEP  = "peep"
)

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

const (
	COW   = "cow"
	WALK  = "walk"
	MOO   = "moo"
	GRASS = "grass"
)

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

const (
	SNAKE   = "snake"
	MICE    = "mice"
	SLITHER = "slither"
	HSSS    = "hsss"
)

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

var animalMap = map[string]Animal{
	BIRD: Bird{
		food:       WORMS,
		locomotion: FLY,
		noise:      PEEP,
	},
	COW: Cow{
		food:       GRASS,
		locomotion: WALK,
		noise:      MOO,
	},
	SNAKE: Snake{
		food:       MICE,
		locomotion: SLITHER,
		noise:      HSSS,
	},
}

var name string

func main() {
	StartProgram()
}

func StartProgram() {
	command := processCommand()

	var name string
	if command == NEW_ANIMAL {
		name = processName()
		processNewAnimal(name)
		StartProgram()
	} else if command == QUERY {
		name = processAnimalQuery()
		processAnimalActivity(name)
		StartProgram()
	} else {
		StartProgram()
	}
}

func getListToPrint(list []string) string {
	return "\"" + strings.Join(list, "\", \"") + "\""
}

func processCommand() string {
	fmt.Printf(ENTER_COMMAND_MSG, commandPrintList)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	checkExit(input)

	command := strings.Fields(input)

	if isInvalid(command[0], commandList) {
		fmt.Printf(INVALID_COMMAND_MSG, command[0], commandPrintList)
		return INVALID_COMMAND
	}

	if strings.ToLower(command[0]) == NEW_ANIMAL {
		return NEW_ANIMAL
	}

	return QUERY
}

func checkExit(input string) {
	if strings.ToLower(input) == "x" {
		fmt.Println(EXIT_MSG)
		os.Exit(0)
	}
}

func processName() string {
	fmt.Println(ENTER_NAME_MSG)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	checkExit(input)

	if len(input) < MIN_ANIMAL_NAME_LENGTH {
		fmt.Printf(INVALID_ANIMAL_NAME_MSG, MIN_ANIMAL_NAME_LENGTH)
		return processName()
	}

	return input
}

func processNewAnimal(name string) {
	fmt.Printf(ENTER_ANIMAL_TYPE_MSG, animalTypePrintList)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	checkExit(input)

	animalTypeField := strings.Fields(input)
	animalType := strings.ToLower(animalTypeField[0])

	if isInvalid(animalType, animalTypeList) {
		fmt.Printf(INVALID_ANIMAL_TYPE_MSG, animalType, animalTypePrintList)
		processNewAnimal(name)
		return
	}

	userAnimals[name] = animalMap[animalType]

	fmt.Println(CREATED_ANIMAL_MSG)
}
func processAnimalQuery() string {
	fmt.Println(ENTER_QUERY_NAME_MSG)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	checkExit(input)

	_, ok := userAnimals[input]
	if ok {
		return input
	}
	fmt.Printf(INVALID_ANIMAL_MSG, input)

	return processAnimalQuery()
}

func processAnimalActivity(name string) {
	fmt.Printf(ENTER_ANIMAL_COMMAND_MSG, animalActivityPrintList)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	checkExit(input)

	animalActivityField := strings.Fields(input)
	animalActivity := strings.ToLower(animalActivityField[0])

	if isInvalid(animalActivity, animalActivityList) {
		fmt.Printf(INVALID_ANIMAL_ACTIVITY_MSG, animalActivity, animalActivityPrintList)
		processAnimalActivity(name)
		return
	}

	if animalActivity == EAT {
		userAnimals[name].Eat()
	} else if animalActivity == MOVE {
		userAnimals[name].Move()
	} else if animalActivity == SPEAK {
		userAnimals[name].Speak()
	}
}

func isInvalid(target string, list []string) bool {
	for _, element := range list {
		if target == element {
			return false
		}
	}
	return true
}
