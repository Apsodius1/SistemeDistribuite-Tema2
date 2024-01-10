package main

import (
	"fmt"
)


// ? Ex1
func isValidForEx1(s string) bool {
	// check if the string has a even number of vowels and a number of consonants divisible by 3
	vowels := 0
	consonants := 0
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		default:
			consonants++
		}
	}
	return vowels%2 == 0 && consonants%3 == 0
}

func countValidStrings(strings []string) int {
	count := 0
	for _, s := range strings {
		if isValidForEx1(s) {
			count++
		}
	}
	return count
}

// ? Ex2
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func countPalindromes(row []string) float64 {
	count := 0
		for _, char := range row {
			if isPalindrome(string(char)) {
				count++
			}
	}
	return float64(count)
}


func main() {
	// ? Ex1
	input1 := [][]string{
		{"aabbb", "ebep", "blablablaa", "hijk", "wsww"},
		{"abba", "eeeppp", "cocor", "ppppppaa", "qwerty","acasq"},
		{"lalala", "lalal", "papapa", "papap"},
	}

	mapResult1 := make(chan int, len(input1))
	reduceResult1 := make(chan int)

	// Map phase
	for _, row := range input1 {
		go func(row []string) {
			mapResult1 <- countValidStrings(row)
		}(row)
	}

	// Reduce phase
	go func() {
		total := 0
		for i := 0; i < len(input1); i++ {
			total += <-mapResult1
		}
		reduceResult1 <- total / len(input1)
	}()

	// Print the result
	result := <-reduceResult1
	fmt.Printf("Average number of strings with an even number of vowels and a number of consonants divisible by 3: %d\n", result)

	// ? Ex2
	input2 := [][]string{
		{"a1551a", "parc", "ana", "minim", "1pcl3"},
		{"calabalac", "tivit", "leu", "zece10", "ploaie", "9ana9"},
		{"lalalal", "tema", "papa", "ger"},
	}

	mapResult2 := make(chan float64, len(input2))
	reduceResult2 := make(chan float64)

	// Map phase
	for _, row := range input2 {
		go func(row []string) {
			mapResult2 <- countPalindromes(row)
		}(row)
	}

	// Reduce phase
	go func() {
		total := 0.0
		for i := 0; i < len(input2); i++ {
			total += <-mapResult2
		}
		reduceResult2 <- total / float64(len(input2))
	}()

	// Print the result
	result2 := <-reduceResult2
	fmt.Printf("Average number of palindromes per row: %.2f\n", result2)
}
