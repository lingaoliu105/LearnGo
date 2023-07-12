package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m
	index2 := 0
	copy(nums1[n:], nums1[:m])
	for i := 0; i < m+n; i++ {
		if nums1[index1] <= nums2[index2] {
			nums1[i] = nums1[index1]
			index1++
		} else {
			nums1[i] = nums2[index2]
			index2++
		}
		if index1 >= m+n {
			copy(nums1[i+1:], nums2[index2:])
			return
		}
		if index2 >= n {
			return
		}
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{4, 0, 0, 0, 0, 0},
		1,
		[]int{1, 2, 3, 5, 6},
		5)
}
