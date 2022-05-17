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
	input_regexp, _ := regexp.Compile("^[+-]{0,1}.+[+-].+")
	if len(input) == 0 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}

	if !(strings.HasPrefix(input, "-") || strings.HasPrefix(input, "+")) {
		input = "+" + input
	}

	if input_regexp.MatchString(input) {
		operandCount := 1
		inputSlice := []rune(input)
		inputLength := len(inputSlice)
		sum := 0
		startIndex := 0
		endIndex := 0
		for i := startIndex + 1; i < inputLength; i++ {
			if inputSlice[i] == '+' || inputSlice[i] == '-' || i == inputLength-1 {
				endIndex = i
				if i == inputLength-1 {
					endIndex = i + 1
				}
				operandCount++

				if operandCount > 2 {
					return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
				}

				convertedValue, err := strconv.Atoi(string(inputSlice[startIndex:endIndex]))
				if err != nil {
					return "", fmt.Errorf("conversion error: %w", err)
				}
				sum += convertedValue
				startIndex = i
			}

		}

		output = strconv.Itoa(sum)
		return output, nil
	}

	return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
}
