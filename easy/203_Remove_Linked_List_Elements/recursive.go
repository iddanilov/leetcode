package main

func recRemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		head = recRemoveElements(head.Next, val)
	} else {
		head.Next = recRemoveElements(head.Next, val)
	}

	return head
}
