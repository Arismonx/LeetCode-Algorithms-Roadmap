package main

import (
	"fmt"
	"runtime/debug"
)

func reverseString(s []byte) {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		s[i], s[lenS-1] = s[lenS-1], s[i]
		lenS--
	}

	fmt.Println(string(s))
}

func init() {
	debug.SetMemoryLimit(1)
}

func main() {
	str := []byte{'h', 'e', 'l', 'l', 'o', 'g'}  // g o l l e h
	str2 := []byte{'h', 'e', 'l', 'l', 'o'}      // o l l e h
	str3 := []byte{'H', 'a', 'n', 'n', 'a', 'h'} // "h","a","n","n","a","H"
	reverseString(str)
	reverseString(str2)
	reverseString(str3)
}
