package main

func plusOne(digits []int) []int {
	return plusOneRecur(digits, len(digits)-1)
}

func plusOneRecur(digits []int, index int) []int {
	if digits[index] != 9 {
		digits[index]++
		return digits
	} else if index == 0 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
		return digits
	} else {
		digits[index] = 0
		return plusOneRecur(digits, index-1)
	}
}
