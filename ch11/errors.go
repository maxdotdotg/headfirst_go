package main

import (
	"fmt"
	"log"
)

type OverheatError float64

// implement Error method to satisfy the error interface
func (o OverheatError) Error() string {
	return fmt.Sprintf("overheating by %0.2f degrees", o)
}

// use OverheatError
func checkTemp(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = checkTemp(121.22, 100.00)
	if err != nil {
		log.Fatal(err)
	}
}
