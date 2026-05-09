package core

import (
	"fmt"
	"go/format"
)

func PrintFormatted(input string) string {
	code, err := formatGoCode(input)
	if err != nil {
		return input
	}
	return code
}

func formatGoCode(input string) (string, error) {
	// Convert the string to a []byte
	source := []byte(input)

	// Format the code using go/format
	formattedSource, err := format.Source(source)
	if err != nil {
		return "", fmt.Errorf("failed to format code: %w", err)
	}

	// Convert the formatted code back to a string
	return string(formattedSource), nil
}
