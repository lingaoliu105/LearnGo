package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tilingRectangle(n int, m int) int {
	ans := max(n, m)

	//initialize rectangle
	rect := make([][]bool, n)
	for i := 0; i < n; i++ {
		rect[i] = make([]bool, m)
	}

	//function to check if square of size k can be placed at x,y
	isAvailable := func(x, y, k int) bool {
		//all slots to be covered must be false (available)
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				} //internal functions are able to access variables from outside
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			return
		}

		if x >= n {
			ans = cnt
			return
		}

		//寻找下一个能放正方形的点位
		if y >= m {
			dfs(x+1, 0, cnt)
			return
		}
		if rect[x][y] {
			dfs(x, y+1, cnt)
			return
		}

		len := min(n-x, m-y)
		for {
			if len <= 0 || isAvailable(x, y, len) {
				break
			}
			len--
		}
		for k := len; k >= 1; k-- {
			fillUp(x, y, k, true)
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}
	}
	dfs(0, 0, 0)
	return ans

}
