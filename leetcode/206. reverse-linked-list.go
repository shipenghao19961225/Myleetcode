package leetcode

// 递归解法
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	// 这里是为了防止最后一个结点和倒数第二个结点形成环
	head.Next = nil
	return res
}

// 迭代法
func reverseList(head *ListNode) *ListNode {
	var p1 *ListNode
	p2 := head
	for p2 != nil {
		tmp := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = tmp
	}
	return p1
}
