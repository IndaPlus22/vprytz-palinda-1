package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// setup hashmap
	words := make(map[string]int)

	// split s by whitespace and iterate over each word
	for _, word := range strings.Fields(s) {
		words[word] += 1
	}

	return words
}

func main() {
	wc.Test(WordCount)
}
