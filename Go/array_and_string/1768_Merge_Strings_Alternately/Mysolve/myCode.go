package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	runes1 := []rune(word1)
	runes2 := []rune(word2)

	len1 := len(runes1)
	len2 := len(runes2)

	mexLength := len1 + len2
	resultRune := make([]rune, 0, mexLength)

	i, j := 0, 0

	for i < len1 && j < len2 {
		resultRune = append(resultRune, runes1[i])
		resultRune = append(resultRune, runes2[j])
		i++
		j++
	}

	for i < len1 {
		resultRune = append(resultRune, runes1[i])
		i++
	}

	for j < len2 {
		resultRune = append(resultRune, runes2[j])
		j++
	}

	return string(resultRune)
}

func main() {
	fmt.Println("output:", mergeAlternately("abc", "pqr")) // apbqcr
	fmt.Println("output:", mergeAlternately("ab", "pqrs")) // apbqrs
}
