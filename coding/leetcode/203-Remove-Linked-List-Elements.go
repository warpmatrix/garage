/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dump := ListNode{ Next : head }
    for pNode, node := &dump, head; node != nil; node = pNode.Next {
        if node.Val == val {
            pNode.Next = node.Next
            // deconstruction function of node
        } else {
            pNode = node
        }
    }
    return dump.Next
}
