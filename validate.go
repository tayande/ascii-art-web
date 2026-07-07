package main

import (
	"fmt"
)
func ValidateInput(text string) (rune, error) {
	for _, char := range text {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("%c is not a valid ascii letter", char)
		}
	}
	return 0, nil
}