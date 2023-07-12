package main

import (
	"fmt"
	"sort"
)

func findPrimePairs(n int) [][]int {
	result := make([][]int, 0)
	var y int
	for x := 0; x <= n/2; x++ {
		y = n - x
		if isPrime(x) && isPrime(y) {
			result = append(result, []int{x, y})
		}
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	longest := 1
	length := 0
	for right := 1; right < len(nums); right++ {
		if nums[right] > threshold {
			length = 0
			continue
		}
		if nums[right]%2 == 0 && length == 0 {
			length = 1
			continue
		}
		if length >= 1 {
			if nums[right]%2 == 0 && nums[right-1]%2 == 1 {
				length++
			} else if nums[right]%2 == 1 && nums[right-1]%2 == 0 {
				length++
			} else {
				length = 0
			}
		}

		if length > longest {
			longest = length
		}
	}
	return longest
}

func continuousSubarrays(nums []int) int64 {

}

func sumImbalanceNumbers(nums []int) int {

}

func getImbalanceNumber(nums []int) int {
	sarr := make([]int, len(nums))
	copy(sarr, nums)
	sort.Ints(sarr)
	cnt := 0
	for index, val := range sarr {
		if val-nums[index] > 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	longestAlternatingSubarray([]int{2, 3, 4, 5}, 4)
	output := make([]int, 0)
	for i := 2; i < 500000; i++ {
		if isPrime(i) {
			output = append(output, i)
		}
	}
	fmt.Println(output)
}
