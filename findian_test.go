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

func TestWrongFirstChar(t *testing.T) {
	reader, writer := setup()

	input := "x"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestWrongMidChar(t *testing.T) {
	reader, writer := setup()

	input := "ix"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestWrongLastChar(t *testing.T) {
	reader, writer := setup()

	input := "iax"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestIncorrectChar1(t *testing.T) {
	reader, writer := setup()

	input := "ihhhhhn"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestIncorrectChar2(t *testing.T) {
	reader, writer := setup()

	input := "ina"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestIncorrectChar3(t *testing.T) {
	reader, writer := setup()

	input := "xian"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestCorrectChar1(t *testing.T) {
	reader, writer := setup()

	input := "ian"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, VALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", VALID_MSG)
	}
}

func TestCorrectChar2(t *testing.T) {
	reader, writer := setup()

	input := "Ian"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, VALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", VALID_MSG)
	}
}

func TestCorrectChar3(t *testing.T) {
	reader, writer := setup()

	input := "iuiygaygn"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, VALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", VALID_MSG)
	}
}

func TestCorrectChar4(t *testing.T) {
	reader, writer := setup()

	input := "I d skd a efju N"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	output := teardownAndGetOutput(reader, writer)

	if !strings.Contains(output, VALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", VALID_MSG)
	}
}
