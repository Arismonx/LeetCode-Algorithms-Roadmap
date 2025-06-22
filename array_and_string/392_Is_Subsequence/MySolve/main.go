package main

import "fmt"

func isSubsequence(s string, t string) bool {
	valid1, valid2 := 0, 0
	for valid1 < len(s) && valid2 < len(t) {
		if s[valid1] == t[valid2] {
			valid1 += 1
		}
		valid2 += 1
	}
	return valid1 == len(s)
}

func main() {
	fmt.Println(isSubsequence("acb", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))

}
