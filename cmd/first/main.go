package main

import (
	"fmt"
	"strings"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// your code goes here
	s = strings.ToLower(s)
	var ret []rune
	cur := 1
	for i := 0; i < len(s); i++ {
		val := rune(s[i])
		if (val >= 'a' && val <= 'z') || (val >= 0 && val <= 9) {
			if cur == 3 {
				cur = 1
				if s[i] >= 'a' && s[i] <= 'z' {
					val = rune(s[i] - 32)
				}
			} else {
				cur = cur + 1
			}
		}
		ret = append(ret, val)
	}
	return string(ret)
}

func main() {
	s := CapitalizeEveryThirdAlphanumericChar("Aspiration.com")
	fmt.Println(s)
}
