package main

func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		val, ok := memo[n]
		if ok {
			return val
		} else {
			memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
			return memo[n]
		}

	}
}

func climbStairsIterative(n int) int {
	a, b, c := 0, 0, 1
	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

func maxProfit(prices []int) int {
	MaxProfit := 0
	l := len(prices)
	if prices[l-2] < prices[l-1] {
		MaxProfit = prices[l-1] - prices[l-2]
	}
	for i := len(prices) - 3; i >= 0; i-- {
		if prices[i] < prices[i+1] {
			MaxProfit += prices[i+1] - prices[i]
		}
	}
	return MaxProfit
}

func maxSubArray(nums []int) int {
	maxSum := 0
	maxSumToEnd := 0
	for _, value := range nums {
		if maxSumToEnd < 0 {
			maxSumToEnd = value
		}
		if maxSumToEnd > maxSum {
			maxSum = maxSumToEnd
		}
	}
	return maxSum
}

//func maxSubArrayDivideNConquer(nums []int, l, r int) (lsum, rsum, msum, isum int) {
//	if l == r {
//		num := nums[l]
//		lsum, rsum = num, num
//		if num > 0 {
//			msum = num
//		} else {
//			msum = 0
//		}
//		isum = num
//		return
//	} else {
//		m := (l + r) / 2
//		l1, l2, l3, l4 := maxSubArrayDivideNConquer(nums, l, m)
//		r1, r2, r3, r4 := maxSubArrayDivideNConquer(nums, m+1, r)
//
//	}
//}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	l0 := nums[0]
	l1 := max(nums[0], nums[1])
	var l2 int
	for i := 2; i < len(nums); i++ {
		l2 = max(l0+nums[i], l1)
		l0 = l1
		l1 = l2
	}
	return l2
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
