// Package datafile reads data samples from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line in a file
// we were hard-coding the size of the array returned
// but not anymore, now we're using slices
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		// return nil (zero value for a slice) instead of bad data
		return nil, err
	}

	// scan the file and convert the string to float64
	// then append it to the numbers slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			// return nil (zero value for a slice) instead of bad data
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}
