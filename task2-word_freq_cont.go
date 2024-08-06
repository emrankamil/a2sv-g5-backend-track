package main

import (
	"fmt"
	"strings"
	"unicode"
)

func WordFreqCounter(word string) map[string]int64{
	words := strings.Fields(word)
	counter := make(map[string]int64)

	for _, word:= range words{
		var clean_word string
		for _, chr:= range word{
			if unicode.IsLetter(chr){
				clean_word += string(chr)
			}
		}
		if clean_word != ""{
			counter[clean_word] += 1
		}
	}

	fmt.Println(counter)

	return counter
}