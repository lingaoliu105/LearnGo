package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val < list2.Val {
			mergeTwoListsRecur(list1.Next, list2, list1)
			return list1
		} else {
			mergeTwoListsRecur(list1, list2.Next, list2)
			return list2

		}
	}
}

func mergeTwoListsRecur(list1, list2, basis *ListNode) {
	if list1 == nil {
		basis.Next = list2
	} else if list2 == nil {
		basis.Next = list1
	} else {
		if list1.Val < list2.Val {
			basis.Next = list1
			mergeTwoListsRecur(list1.Next, list2, list1)
		} else {
			basis.Next = list2
			mergeTwoListsRecur(list1, list2.Next, list2)
		}
	}
}
