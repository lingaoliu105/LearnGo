package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		return describe(countAndSay(n - 1))
	}
}

func describe(s string) string {
	l := len(s)
	if l == 1 {
		return "1" + s
	}
	prev := s[0]
	cnt := 1
	result := make([]byte, 0)
	for i := 1; i < l; i++ {
		if s[i] != prev {
			result = strconv.AppendInt(result, int64(cnt), 10)
			result = strconv.AppendInt(result, int64(prev-48), 10)
			prev = s[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	result = strconv.AppendInt(result, int64(cnt), 10)
	result = strconv.AppendInt(result, int64(prev-48), 10)
	return string(result)
}

func main() {
	fmt.Println(countAndSay(4))
}
