package main

import (
	"fmt"
	"strings"
	"unicode"
)

func PalindromeCheck(word string){
	rstr := []rune(strings.ToLower(word))
	i, j := 0, len(word) - 1
	var isPalindrome bool
	
	for i < j{
		if  ! unicode.IsLetter(rstr[i]){
			i += 1
		}else if ! unicode.IsLetter(rstr[j]){
			j -= 1
		}else if word[i] != word[j]{
			isPalindrome = false
			break
		}

		i += 1
		j -= 1
	}

	isPalindrome = true

	if isPalindrome{
		fmt.Println("the word-", word, "-is palindrome")
	}else{
		fmt.Println("the word-", word, "-is not palindrome")
	}
}