package main

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	temp := make([]*TreeNode, 0)
	var node *TreeNode
	cnt := 0
	for len(q) > 0 {
		if len(result) <= cnt {
			result = append(result, []int{})
		}
		for len(q) > 0 {
			node = q[0]
			q = q[1:]
			result[cnt] = append(result[cnt], node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		cnt++
		q, temp = temp, q
	}
	return result
}
