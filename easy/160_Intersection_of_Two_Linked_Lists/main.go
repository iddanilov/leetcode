package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	list := make(map[*ListNode]interface{})

	for headA != nil {
		list[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := list[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func main() {
	fmt.Println(getIntersectionNode(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		}))
}
