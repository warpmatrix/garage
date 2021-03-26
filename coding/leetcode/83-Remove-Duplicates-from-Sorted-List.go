func deleteDuplicates(head *ListNode) *ListNode {
    // val range [-100, 100]
	dump := ListNode{-200, head}
	for node, prev := head, &dump; node != nil; node = node.Next {
		if node.Next != nil && node.Val == node.Next.Val {
            prev.Next = node.Next
		} else {
            prev = node
        }
	}
	return dump.Next
}
