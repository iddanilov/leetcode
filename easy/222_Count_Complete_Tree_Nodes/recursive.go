package main

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 1
	if root.Left != nil {
		result += countNodes(root.Left)
	}
	if root.Right != nil {
		result += countNodes(root.Right)
	}

	return result
}
