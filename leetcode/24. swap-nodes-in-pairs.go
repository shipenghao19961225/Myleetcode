package leetcode

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	tmp := dummy
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next
		node1.Next = node2.Next
		tmp.Next = node2
		node2.Next = node1
		tmp = node1
	}
	return dummy.Next
}
