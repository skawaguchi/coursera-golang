/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user to enter values for
- acceleration
- initial velocity
- initial displacement

- Then the program should prompt the user to enter a value for time
- the program should compute the displacement after the entered time.

- You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so.
- GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
- The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

EXAMPLE
Let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Validated answers against: https://www.calculatorsoup.com/calculators/physics/displacement_v_a_t.php
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var ACCELERATION_MSG = "Please enter a value for `Acceleration`:"
var VELOCITY_MSG = "Please enter a value for `Initial Velocity`:"
var DISPLACEMENT_MSG = "Please enter a value for `Initial Displacement`:"
var TIME_MSG = "Please enter a value for `Time`:"
var FINAL_OUTPUT_MSG = "The `Final Displacement` is:"

var INPUT_ERROR = "You did not type a float as a valid entry for %q. Please try again."

type Inputs struct {
	acceleration float64
	displacement float64
	velocity     float64
}

type dataEnum int

const (
	ACCELERATION dataEnum = iota
	DISPLACEMENT dataEnum = iota
	TIME         dataEnum = iota
	VELOCITY     dataEnum = iota
)

func (de dataEnum) String() string {
	switch de {
	case ACCELERATION:
		return "Acceleration"
	case DISPLACEMENT:
		return "Displacement"
	case TIME:
		return "Time"
	default:
		return "Velocity"
	}
}

func main() {
	inputs := getInitialInputs()
	genFn := GenDisplaceFn(inputs.acceleration, inputs.velocity, inputs.displacement)

	accelerationNum := getInput(TIME, TIME_MSG)

	fmt.Println(FINAL_OUTPUT_MSG, genFn(accelerationNum))
}

func getInitialInputs() Inputs {
	accelerationNum := getInput(ACCELERATION, ACCELERATION_MSG)
	velocityNum := getInput(VELOCITY, VELOCITY_MSG)
	displacementNum := getInput(DISPLACEMENT, DISPLACEMENT_MSG)

	inputs := Inputs{
		acceleration: accelerationNum,
		displacement: displacementNum,
		velocity:     velocityNum,
	}

	return inputs
}

func getInput(inputType dataEnum, msg string) float64 {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(msg)

	scanner.Scan()

	input := scanner.Text()

	num, strConvErr := strconv.ParseFloat(input, 64)

	if strConvErr != nil {
		log.Fatalf(INPUT_ERROR, inputType)
	}

	return num
}

func GenDisplaceFn(
	acceleration float64,
	initialVelocity float64,
	initialDisplacement float64,
) func(float64) float64 {
	return func(time float64) float64 {
		return (math.Pow(time, 2) * acceleration / 2) + initialVelocity*time + initialDisplacement
	}
}
