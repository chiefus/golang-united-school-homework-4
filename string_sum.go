package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	input_regexp, _ := regexp.Compile("^[+-].+[+-].+")
	if len(input) == 0 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}

	if !(strings.HasPrefix(input, "-") || strings.HasPrefix(input, "+")) {
		input = "+" + input
	}

	if input_regexp.MatchString(input) {
		operandCount := 0
		inputSlice := []rune(input)
		inputLength := len(inputSlice)
		sum := 0

		for i := 0; i < inputLength; i++ {
			if inputSlice[i] == '+' || inputSlice[i] == '-' {
				operandCount++
				startIndex := i
				endIndex := startIndex
				for j := startIndex + 1; j < inputLength; j++ {
					endIndex = j
					i = j - 1
					if inputSlice[j] == '+' || inputSlice[j] == '-' {
						break
					}
					if j == inputLength-1 {
						endIndex = inputLength
					}
				}

				if operandCount > 2 {
					return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
				}
				fmt.Println(string(inputSlice[startIndex:endIndex]))
				convertedValue, err := strconv.Atoi(string(inputSlice[startIndex:endIndex]))
				if err != nil {
					return "", fmt.Errorf("conversion error: %w", err)
				}
				sum += convertedValue
			}

		}

		output = strconv.Itoa(sum)
		return output, nil
	}

	return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
}
