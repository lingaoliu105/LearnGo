package main

import (
	"fmt"
	"unsafe"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// var beforePrev uintptr
	if head == nil || head.Next == nil {
		return false
	}
	var prev uintptr = uintptr(unsafe.Pointer(head))
	prev2 := prev

	var curr uintptr
	head = head.Next
	for head != nil {
		curr = uintptr(unsafe.Pointer(head))
		if prev|curr == prev && prev2&curr == prev2 {

			return true
		}
		// beforePrev = prev
		prev = curr | prev
		prev2 = curr & prev2

		head = head.Next
	}
	return false
}
func makeCycledLinkedList(arr []int) *ListNode {
	head := ListNode{arr[0], nil}
	headPtr := &head
	currPtr := headPtr
	for _, num := range arr[1:] {
		newNode := &ListNode{num, headPtr}
		currPtr.Next = newNode
		currPtr = newNode
	}
	return headPtr
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
	head := makeLinkedList([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5})
	fmt.Println(hasCycle(head))
}
