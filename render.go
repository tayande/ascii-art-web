package main


import (
	"strings"
)
func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for row := 0; row < 8; row++ {
		var builder strings.Builder
		for _, char := range text {
			builder.WriteString(banner[char][row])
		}
		result[row] = builder.String()
	}
	return result
}