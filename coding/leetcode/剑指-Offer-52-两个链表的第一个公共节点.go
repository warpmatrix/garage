/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil { return nil }
    ptrA, ptrB := headA, headB
    for ptrA != ptrB {
        ptrA, ptrB = ptrA.Next, ptrB.Next
        if ptrA == nil && ptrB != nil {
            ptrA = headB
        } else if ptrA != nil && ptrB == nil {
            ptrB = headA
        }
    }
    return ptrA
}
