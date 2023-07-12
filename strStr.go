package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if isPrefix(haystack[i+1:], needle[1:]) {
				return i
			}
		}
	}
	return -1
}

func isPrefix(haystack string, needle string) bool {
	if len(needle) > len(haystack) {
		return false
	}
	if len(haystack) > 0 && len(needle) > 0 && haystack[0] != needle[0] {
		return false
	} else if len(needle) <= 0 {
		return true
	}
	return isPrefix(haystack[1:], needle[1:])
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
