package main

func isPalindrome(head *ListNode) bool {
	arrayList := make([]int, 0)
	buildArray(head, &arrayList)
	return sliceIsPalindrome(arrayList)

}

func buildArray(head *ListNode, sliPtr *[]int) {
	*sliPtr = append(*sliPtr, head.Val)
	buildArray(head.Next, sliPtr)
}

func sliceIsPalindrome(sl []int) bool {
	if len(sl) == 0 || len(sl) == 1 {
		return true
	} else if sl[0] != sl[len(sl)-1] {
		return false
	} else {
		return sliceIsPalindrome(sl[1 : len(sl)-1])
	}
}

func isPalindromeRecur(head *ListNode) bool {
	outerPtr := &head
	var isPal *bool = new(bool) //通过new()实现malloc
	*isPal = true
	var scanInner func(*ListNode)
	scanInner = func(inner *ListNode) {
		if inner == nil {
			return
		} else {
			scanInner(inner.Next)
			if (*outerPtr).Val != inner.Val {
				*isPal = false
				return
			}
			*outerPtr = (*outerPtr).Next
		}
	}
	scanInner(head)
	return *isPal
}
