package main

import (
	"fmt"
	"strings"
)

func WordFrequency(inputs string) map[string]int {

	words := strings.Fields(inputs)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func main() {
	testWords := "Golang is one of the very very best of the best programming language"
	result := WordFrequency(testWords)
	fmt.Print(result)
}
