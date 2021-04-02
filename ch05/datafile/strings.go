// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line in a file.
func GetStrings(fileName string) ([]string, error) {
	// insantiate empty slice for lines in the file
	var lines []string

	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// read each line in the file, and append it to lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return lines, nil
}
