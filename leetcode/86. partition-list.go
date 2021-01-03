package leetcode

// 画图，加哑结点即可
func partition(head *ListNode, x int) *ListNode {
	lessYummy, moreYummy := new(ListNode), new(ListNode)
	lessP, moreP := lessYummy, moreYummy
	for head != nil {
		if head.Val < x {
			lessP.Next = head
			lessP = lessP.Next
		} else {
			moreP.Next = head
			moreP = moreP.Next
		}
		head = head.Next
	}
	lessP.Next = moreYummy.Next
	moreP.Next = nil
	return lessYummy.Next
}
