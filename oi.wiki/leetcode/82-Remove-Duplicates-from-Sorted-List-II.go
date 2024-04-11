/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    // val range [-100, 100]
	dump := ListNode{-200, head}
	for node, prev := head, &dump; node != nil; {
		if node.Next != nil && node.Val == node.Next.Val {
			tarVal := node.Val
			iter := node.Next.Next
			for iter != nil && iter.Val == tarVal {
				iter = iter.Next
			}
			prev.Next = iter
			node = iter
		} else {
			prev, node = prev.Next, node.Next
		}
	}
	return dump.Next
}
