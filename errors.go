package main

import (
	"errors"
	"fmt"
	"strconv"
)

func convert(x string) {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(wentWrong)
	} else {
		fmt.Println(i)
	}
}

// custom error
type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", de.dividend)
}

//error.New

var wentWrong error = errors.New("Something went wrong")
