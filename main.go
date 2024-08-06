package main

import (
	"fmt"
	"os"
	"bufio"
)

// run this code as:
// go run Task-1-student_grade.go Task-2.1-word_freq_cont.go Task-2.2-palindrome_check.go main.go

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose which task to run:")
	fmt.Println("1. Task 1: Student Grade Calculation")
	fmt.Println("2. Task 2.1: Word Frequency Counter")
	fmt.Println("3. Task 2.2: Palindrome Checker")

	fmt.Print("Enter your choice (1-3): ")
	choice, _ := reader.ReadString('\n')

	switch choice {
	case "1\n":
		GradeCalculator()
	case "2\n":
		WordFreqCounter("this code will count frequency of each words in this sentence")
	case "3\n":
		PalindromeCheck("word")
	default:
		fmt.Println("Invalid choice")
	}
}
