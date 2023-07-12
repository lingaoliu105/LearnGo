package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head
	var rightBeforeLeft *ListNode
	cnt := 0
	for right.Next != nil {
		right = right.Next
		cnt++
		if cnt >= n {
			rightBeforeLeft = left
			left = left.Next
		}
	}
	if left == head {
		return left.Next
	}
	rightBeforeLeft.Next = left.Next
	return head
}

func makeLinkedList(arr []int) *ListNode {
	head := ListNode{arr[0], nil}
	headPtr := &head
	currPtr := headPtr
	for _, num := range arr[1:] {
		newNode := &ListNode{num, nil}
		currPtr.Next = newNode
		currPtr = newNode
	}
	return headPtr
}
func main() {
	input := makeLinkedList([]int{1, 2, 3, 4, 5})
	removeNthFromEnd(input, 2)
}
