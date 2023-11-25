package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l11 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l12 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	addTwoNumbers(&l11, &l12)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	switch {
	case l1.Next == nil && l2.Next == nil:
		return nil
	case l1 == nil:
		result = l1
	case l2 == nil:
		result = l2
	default:
		result = addTwoNumbers(l1, l2)
	}
	return result
}

func sum(l1 *ListNode, n1, n2 int) {
	if n1 +new
}
