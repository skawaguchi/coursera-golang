package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
// https://stackoverflow.com/questions/47281081/how-to-test-method-that-prints-to-console-in-golang/47281683

var origStdOut = os.Stdout

func setup() (reader io.Reader, writer os.File) {
	readerVal, writerVal, _ := os.Pipe()
	os.Stdout = writerVal

	return io.Reader(readerVal), *writerVal
}

func teardownAndGetOutput(reader io.Reader, writer os.File) (output string) {
	writer.Close()

	outputVal, _ := ioutil.ReadAll(reader)

	strOutput := string(outputVal)

	os.Stdout = origStdOut

	return strOutput
}

func assertInput(t *testing.T, testInput string, expectedMessage string) {
	reader, writer := setup()

	input := testInput
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, expectedMessage) {
		t.Errorf("output did NOT contain expected message %q", expectedMessage)
	}
}

func TestWrongFirstChar(t *testing.T) {
	assertInput(t, "x", INVALID_MSG)
}

func TestWrongMidChar(t *testing.T) {
	assertInput(t, "ix", INVALID_MSG)
}

func TestWrongLastChar(t *testing.T) {
	assertInput(t, "iax", INVALID_MSG)
}

func TestIncorrectChar1(t *testing.T) {
	assertInput(t, "ihhhhhn", INVALID_MSG)
}

func TestIncorrectChar2(t *testing.T) {
	assertInput(t, "ina", INVALID_MSG)
}

func TestIncorrectChar3(t *testing.T) {
	assertInput(t, "xian", INVALID_MSG)
}

func TestCorrectChar1(t *testing.T) {
	assertInput(t, "ian", VALID_MSG)
}

func TestCorrectChar2(t *testing.T) {
	assertInput(t, "Ian", VALID_MSG)
}

func TestCorrectChar3(t *testing.T) {
	assertInput(t, "iuiygaygn", VALID_MSG)
}

func TestCorrectChar4(t *testing.T) {
	assertInput(t, "I d skd a efju N", VALID_MSG)
}
