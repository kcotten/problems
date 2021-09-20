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
