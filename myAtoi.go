package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	sign := 1
	if len(s) <= 0 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	result := 0
	for len(s) > 0 && s[0] >= 48 && s[0] <= 57 {
		result *= 10
		result += int(s[0] - 48)
		s = s[1:]
		if sign == -1 && result*sign <= math.MinInt32 {
			return math.MinInt32
		} else if sign == 1 && result >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return result * sign
}

func main() {
	myAtoi("-2147483647")
}
