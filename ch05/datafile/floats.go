// Package datafile reads data samples from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line in a file
// in this case, we're hard-coding the size of the array returned
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	// store index in i
	i := 0

	// scan the file and convert the string to float64
	// then add it to the numbers array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	return numbers, nil
}
