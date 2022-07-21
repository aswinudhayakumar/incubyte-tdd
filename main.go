package main

import (
	"errors"
	"fmt"
)

var NegativeYearError = errors.New("Year cannot be nagative")

func IsLeapYear(year int) (bool, error) {
	if year < 0 {
		return false, NegativeYearError
	}
	if year%4000 == 0 {
		return false, nil
	}
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true, nil
	}
	return false, nil
}

func main() {
	fmt.Println("Hello world")
}
