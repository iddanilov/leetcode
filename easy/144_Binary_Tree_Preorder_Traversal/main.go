package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(preorderTraversal(&TreeNode{}))
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return []int{}
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
