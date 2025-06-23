package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	m := map[int]string{}
	result := 0
	for i := range jewels {
		m[i] = string(jewels[i])
	}
	for _, v := range m {
		for j := range stones {
			if v == string(stones[j]) {
				result += 1
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
