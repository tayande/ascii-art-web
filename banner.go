package main

import (
	"fmt"
	"os"
	"strings"
)
func LoadBanner(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("banner file does not exist")
	}
	info := string(data)
	if info == "" {
		return nil, fmt.Errorf("banner file is empty")
	}
	lines := strings.Split(info, "\n")
	if len(lines) < 855 {
		return nil, fmt.Errorf("banner file does not have enough characters")
	}
	charMap := map[rune][]string{}
	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9 + 1
		end := start + 8
		charMap[rune(i)] = lines[start:end]
	}
	return charMap, nil
}