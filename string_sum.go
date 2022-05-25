package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	s := strings.ReplaceAll(input, " ", "")

	c := 0
	for _, val := range s[1:] {
		if val == '+' || val == '-' {
			c++
		}
	}

	if c != 1 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	idx := 0

	for i, val := range s {
		if (val == '-' || val == '+') && i != 0 {
			idx = i
		}
	}

	num1, err := strconv.Atoi(s[:idx])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	num2, err := strconv.Atoi(s[idx+1:])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if s[idx] == '+' {
		output = strconv.Itoa(num1 + num2)
	} else {
		output = strconv.Itoa(num1 - num2)
	}

	return output, nil

}
