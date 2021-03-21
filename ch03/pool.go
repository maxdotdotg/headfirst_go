package main

import (
	"errors"
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) {
	// custom error msg
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

func main() {
	quotient, err := divide(5.6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%02.f\n", quotient)
	}
}
