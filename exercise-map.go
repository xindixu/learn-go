package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts_by_word := make(map[string]int)

	for _, word := range words {
		counts_by_word[word] += 1
	}

	return counts_by_word
}

func main() {
	wc.Test(WordCount)
}
