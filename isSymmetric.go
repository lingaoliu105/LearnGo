package main

func isSymmetric(root *TreeNode) bool {
	flipTree(root.Right)
	return isSame(root.Left, root.Right)
}
func flipTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	flipTree(root.Right)
	flipTree(root.Left)
}

func isSame(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	} else {
		return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
	}
}
