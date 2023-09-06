package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	word_map := make(map[string]int)

	for _, word := range words {
		_, exist := word_map[word]
		if !exist {
			word_map[word] = 1
		} else {
			word_map[word]++
		}
	}
	return word_map
}

func main() {
	wc.Test(WordCount)
}
