package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	isValid, _, _ := isValidBSTRecur(root)
	return isValid
}

func isValidBSTRecur(root *TreeNode) (isValid bool, min, max int) {
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	} else if root.Right == nil {
		childIsValid, childMin, childMax := isValidBSTRecur(root.Left)
		return childIsValid && childMax < root.Val, childMin, root.Val
	} else if root.Left == nil {
		childIsValid, childMin, childMax := isValidBSTRecur(root.Right)
		return childIsValid && childMin > root.Val, root.Val, childMax
	} else {
		childIsValid, childMin, childMax := isValidBSTRecur(root.Left)
		if !childIsValid {
			return false, 0, 0
		}
		childIsValid2, childMin2, childMax2 := isValidBSTRecur(root.Right)
		return childIsValid && childIsValid2 && childMax < root.Val && childMin2 > root.Val, childMin, childMax2
	}
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
func main() {
	root := makeTree([]int{2, 1, 3})
	isValidBST(root)
}
