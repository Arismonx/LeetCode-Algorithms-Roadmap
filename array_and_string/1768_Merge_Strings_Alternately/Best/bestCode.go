package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	lettersToMerge := make([]string, (len(word1)+len(word2))+100)

	var merged string

	word2Count := 0
	for i := range lettersToMerge {
		if word2Count+1 > len(word2) {
			break
		}

		if i%2 == 0 {
			lettersToMerge[i+1] = string([]byte(word2)[word2Count])
			word2Count++
		}
	}

	word1Count := 0
	for i, v := range lettersToMerge {
		if word1Count+1 > len(word1) {
			break
		}

		if len(v) == 0 {
			lettersToMerge[i] = string([]byte(word1)[word1Count])
			word1Count++
		}
	}

	for _, v := range lettersToMerge {
		merged += v
	}

	return merged
}

func main() {
	fmt.Println("output:", mergeAlternately("abc", "pqr")) // apbqcr
	fmt.Println("output:", mergeAlternately("ab", "pqrs")) // apbqrs
}
