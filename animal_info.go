package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var INTRO_MSG = "Please type in an animal TYPE and an ACTIVITY, separated by a SPACE.\nThe choices for type are: %s. The choices for activity are %s. (ex. `%s %s`). \nType `x` to EXIT.\n"
var INVALID_INPUT_MSG = "You must enter %d inputs. Try again.\n"
var INVALID_ANIMAL_MSG = "You typed in %q as the animal type. This is invalid.\n"
var INVALID_ACTIVITY_MSG = "You typed in %q as the animal activity. This is invalid.\n"
var EXIT_MSG = "Bye Bye!"
var EXPECTED_INPUTS = 2

type Animal struct {
	food       string
	locomotion string
	noise      string
}

const (
	EAT   = "eat"
	MOVE  = "move"
	SPEAK = "speak"
)

var animalTypesList = []string{COW, BIRD, SNAKE}
var animalActivitiesList = []string{EAT, MOVE, SPEAK}

const (
	BIRD  = "bird"
	FLY   = "fly"
	WORMS = "worms"
	PEEP  = "peep"
)

var bird = Animal{
	food:       WORMS,
	locomotion: FLY,
	noise:      PEEP,
}

const (
	COW   = "cow"
	WALK  = "walk"
	MOO   = "moo"
	GRASS = "grass"
)

var cow = Animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}

const (
	SNAKE   = "snake"
	MICE    = "mice"
	SLITHER = "slither"
	HSSS    = "hsss"
)

var snake = Animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}
var animalMap = map[string]Animal{
	BIRD:  bird,
	COW:   cow,
	SNAKE: snake,
}

func main() {

	animalTypes := getListToPrint(animalTypesList)
	animalActivities := getListToPrint(animalActivitiesList)

	fmt.Printf(INTRO_MSG, animalTypes, animalActivities, animalTypesList[0], animalActivitiesList[0])

	processInput()
}

func getListToPrint(list []string) string {
	return strings.Join(list, ", ")
}

func processInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "x" {
			fmt.Println(EXIT_MSG)
			break
		}

		fields := strings.Fields(input)

		if len(fields) != EXPECTED_INPUTS {
			fmt.Printf(INVALID_INPUT_MSG, EXPECTED_INPUTS)
			continue
		}

		animalType, animalActivity := fields[0], fields[1]

		if isInvalid(animalType, animalTypesList) {
			fmt.Printf(INVALID_ANIMAL_MSG, animalType)
			continue
		} else if isInvalid(animalActivity, animalActivitiesList) {
			fmt.Printf(INVALID_ACTIVITY_MSG, animalActivity)
			continue
		}

		var animal Animal = animalMap[animalType]

		if animalActivity == EAT {
			animal.Eat()
		} else if animalActivity == MOVE {
			animal.Move()
		} else if animalActivity == SPEAK {
			animal.Speak()
		}

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

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
