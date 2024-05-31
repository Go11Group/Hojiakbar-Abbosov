package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	hashMap := make(map[rune]bool)
	for _, char := range allowed {
		hashMap[char] = true
	}

	n := 0

	for _, word := range words {
		checker := true
		for _, char := range word {
			if !hashMap[char] {
				checker = false
				break
			}
		}
		if checker {
			n++
		}
	}

	return n
}

func main() {
	allowed1, words1 := "ab", []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println("Example 1 Output:", countConsistentStrings(allowed1, words1)) 

	allowed2, words2 := "abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println("Example 2 Output:", countConsistentStrings(allowed2, words2)) 

	allowed3, words3 := "cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println("Example 3 Output:", countConsistentStrings(allowed3, words3)) 
}
