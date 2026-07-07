package main

import (
	"strings"
)
func GenerateArt(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}
	parts := SplitInput(text)
	var result strings.Builder
 
	for i, part := range parts {
		if part == "" {
			if i < len(parts) - 1 {
				result.WriteString("\n")
			} else if part == "" && i == len(parts) - 1 && parts[i-1] != "" {
				result.WriteString("\n")
			}
		} else {
			rows := RenderLine(part, banner)
			for _, row := range rows {
				result.WriteString(row + "\n")
			}
		}
	}
	return result.String()
}