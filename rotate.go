package main

import "fmt"

func rotate(nums []int, k int) {
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func rotate2(nums []int, k int) {
	n := len(nums)
	iter := gcd(k, n)
	for i := 0; i < iter; i++ {
		cnt := 0
		old := nums[i]
		temp := 0
		for (i+cnt*k)%n != i {
			temp = nums[(i+cnt*k)%n]
			nums[(i+cnt*k)%n] = old
			old = temp
			cnt++
		}
		nums[i] = old
	}
}
func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(arr []int, start int, end int) {
	if start == end {
		return
	}
	arr[start], arr[end] = arr[end], arr[start]
	if start == end-1 {

		return
	}
	reverse(arr, start+1, end-1)
}

func swap(arr []int, op1 int, op2 int) {

}
func main() {
	bools := [9]bool{}
	fmt.Println(bools)
	lookUp := make(map[int]int, 10)
	lookUp[4] = 9
	value, ok := lookUp[2]
	value2 := lookUp[1]
	fmt.Println(lookUp, value, ok, value2)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	fmt.Println(nums)
}
