package main

func recursivePostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := recursivePostorderTraversal(root.Left)
	right := recursivePostorderTraversal(root.Right)
	return append(left, append(right, root.Val)...)
}
