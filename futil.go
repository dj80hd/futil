// Package futil provides convenience methods for reading and
// writing file content as strings or slices of strings (lines)
package futil

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// GetLines returns the lines of the given file
func GetLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// GetContent returns the contents of a file as a string
func GetContent(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// SetContent sets the contents of a given file
// If the file does not exist, it will be created
func SetContent(filename string, contents string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, contents)
	return w.Flush()
}

// SetLines sets the lines of a given file
// If the file does not exist, it will be created
func SetLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
