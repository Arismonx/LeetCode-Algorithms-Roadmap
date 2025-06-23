package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	pref := strs[0]
	pref_len := len(pref)

	for i := 1; i < len(strs); i++ {
		s := strs[i]
		for pref_len > len(s) || pref != s[0:pref_len] {
			pref_len--
			if pref_len == 0 {
				return ""
			}
			pref = pref[0:pref_len]
		}
	}
	return pref
}

func main() {
	str := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))
}
