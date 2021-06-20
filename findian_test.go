package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestWrongFirstChar(t *testing.T) {
	// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
	// https://stackoverflow.com/questions/47281081/how-to-test-method-that-prints-to-console-in-golang/47281683

	origStdOut := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	input := "x"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	writer.Close()

	output, _ := ioutil.ReadAll(reader)

	strOutput := string(output)

	os.Stdout = origStdOut

	if !strings.Contains(strOutput, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestWrongMidChar(t *testing.T) {
	origStdOut := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	input := "ix"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	writer.Close()

	output, _ := ioutil.ReadAll(reader)

	strOutput := string(output)

	os.Stdout = origStdOut

	if !strings.Contains(strOutput, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestWrongLastChar(t *testing.T) {
	origStdOut := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	input := "iax"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	writer.Close()

	output, _ := ioutil.ReadAll(reader)

	strOutput := string(output)

	os.Stdout = origStdOut

	if !strings.Contains(strOutput, INVALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", INVALID_MSG)
	}
}

func TestCorrectChar(t *testing.T) {
	origStdOut := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	input := "ian"
	scanner := bufio.NewScanner(strings.NewReader(input))

	Find(scanner)

	writer.Close()

	output, _ := ioutil.ReadAll(reader)

	strOutput := string(output)

	os.Stdout = origStdOut

	if !strings.Contains(strOutput, VALID_MSG) {
		t.Errorf("output did NOT contain expected message %q", VALID_MSG)
	}
}
