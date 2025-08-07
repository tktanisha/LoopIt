package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func IntConversion(input string) int {

	// Optional: Trim spaces
	input = strings.TrimSpace(input)

	// Try to convert input to int
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return 0
	}

	return number
}
