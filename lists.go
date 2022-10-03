package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = new(ListNode)
	curr := l3
	for l1 != nil || l2 != nil {
		var t *ListNode
		switch {
		case l1 == nil:
			t = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next
		case l2 == nil:
			t = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		case l1.Val < l2.Val:
			t = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		case l2.Val < l1.Val:
			t = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next
		case l1.Val == l2.Val:
			t = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		}
		curr.Next = t
		curr = curr.Next
	}
	return l3
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode
	curr = head
	for curr != nil {
		t := curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}
	return prev
}

func deleteDuplicates(head *ListNode) *ListNode {
	l := head
	if head == nil || head.Next == nil {
		return head
	} else {
		for l.Next != nil {
			if l.Next.Val == l.Val {
				l.Next = l.Next.Next
			} else {
				l = l.Next
			}
		}
	}
	return head
}
