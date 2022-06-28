package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words :=strings.Fields(s)
	cnt_words := make(map[string]int)
	
	for i := 0; i < len(words); i++{
		cnt_words[words[i]]++
	}
	return cnt_words
}
func main() {
	wc.Test(WordCount)
}
