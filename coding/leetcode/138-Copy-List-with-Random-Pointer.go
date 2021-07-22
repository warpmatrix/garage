/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    m := map[*Node]*Node{}
    dump := &Node{}
    for ptr, nptr := head, dump; ptr != nil; ptr, nptr = ptr.Next, nptr.Next {
        nptr.Next = &Node{ Val: ptr.Val, Next: ptr.Next }
        m[ptr] = nptr.Next
    }
    for ptr, nptr := head, dump.Next; ptr != nil; ptr, nptr = ptr.Next, nptr.Next {
        nptr.Random = m[ptr.Random]
    }
    return dump.Next
}
