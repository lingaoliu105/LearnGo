package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	//root := makeTree(nums)
	//var inOrderTraversal func(*TreeNode)
	//inOrderTraversal = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	} else {
	//		inOrderTraversal(root.Left)
	//		root.Val = nums[0]
	//		nums = nums[1:]
	//		inOrderTraversal(root.Right)
	//	}
	//}
	//inOrderTraversal(root)
	//return root

	//官方题解做法：
	return helper(nums, 0, len(nums)-1)
}
func makeTree(array []int) *TreeNode {
	q := make([]*TreeNode, 0)
	root := &TreeNode{array[0], nil, nil}
	array = array[1:]
	q = append(q, root)
	var base *TreeNode
	for len(q) > 0 {
		base = q[0]
		q = q[1:]
		if len(array) <= 0 {
			break
		}
		newVal := array[0]
		array = array[1:]
		base.Left = &TreeNode{newVal, nil, nil}
		q = append(q, base.Left)
		if len(array) <= 0 {
			break
		}
		newVal = array[0]
		array = array[1:]
		base.Right = &TreeNode{newVal, nil, nil}
		q = append(q, base.Right)
	}
	return root
}

// 官方题解的做法
func helper(array []int, left, right int) *TreeNode { //left 包含，right不包含
	if left > right {
		return nil
	} else {
		midIndex := (left + right) / 2
		mid := array[midIndex]
		midNodePtr := &TreeNode{mid, nil, nil}
		midNodePtr.Left = helper(array, left, midIndex-1)
		midNodePtr.Right = helper(array, midIndex+1, right)
		return midNodePtr
	}
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
