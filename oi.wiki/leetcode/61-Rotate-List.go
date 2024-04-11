/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    iter := head
    len := 0
    for i := 0; i < k && iter != nil; i++ {
        iter = iter.Next
        len++
    }
    if iter == nil {
        k = k % len
        if k == 0 {
            return head
        }
        iter = head
        for i := 0; i < k; i++ {
            iter = iter.Next
        }
    }
    prev := head
    for ; iter.Next != nil; iter = iter.Next {
        prev = prev.Next
    }
    iter.Next = head
    newHead := prev.Next
    prev.Next = nil
    head = newHead
    return head
}
