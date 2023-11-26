package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	pre, tmp := head, head

	for tmp != nil {
		if tmp.Val == val {
			pre.Next = tmp.Next
		} else {
			pre = tmp
		}

		tmp = tmp.Next
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}
