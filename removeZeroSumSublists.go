package main

import "fmt"

// Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeZeroSumSublists(head *ListNode) *ListNode {
	curr := head
	for {
		cnt := 0
		sums := make([]ListNode, 0)
		flag := false
		for curr != nil {
			sums = append(sums, ListNode{0, curr})
			for index, _ := range sums {
				sums[index].Val += curr.Val
				if sums[index].Val == 0 {
					flag = true
					if index >= 1 {
						sums[index-1].Next.Next = curr.Next
					} else {
						head = curr.Next
					}
					removeZeroSumSublists(curr.Next)
				}
			}
			curr = curr.Next
			cnt++
		}

		if !flag { //if did not find any more zero sum, return.
			return head
		}
	}
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
	head := makeLinkedList([]int{1, 2, -3, 3, 1})
	fmt.Println(head)
	result := removeZeroSumSublists(head)
	fmt.Println(result)
}
