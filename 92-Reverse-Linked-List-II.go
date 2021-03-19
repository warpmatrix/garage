/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var begin *ListNode
    for i := 1; i < left; i++ {
        if i == 1 {
            begin = head
        } else {
            begin = begin.Next
        }
    }
    front := head
    if begin != nil {
        front = begin.Next
    }
    prev := front
    node := prev.Next
    for i := left + 1; i <= right; i++ {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
    }
    front.Next = node
    if begin != nil {
        begin.Next = prev
    } else {
        head = prev
    }
    return head
}