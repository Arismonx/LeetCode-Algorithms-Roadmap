package main

import "fmt"

func romanToInt(s string) int {
	romam := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romam[s[i]] < romam[s[i+1]] {
			result -= romam[s[i]]
		} else {
			result += romam[s[i]]
		}
	}

	return result
}

func main() {
	fmt.Println("s:", romanToInt("III"))
	fmt.Println("s:", romanToInt("LVIII"))
	fmt.Println("s:", romanToInt("MCMXCIV"))
}
