package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	for _, s := range strings.Fields(s) {
		wordCount[s]++
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
