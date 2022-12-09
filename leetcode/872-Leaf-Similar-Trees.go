/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    cur1, cur2 := root1, root2
    stack1, stack2 := []*TreeNode{}, []*TreeNode{}
    leaf1, cur1, stack1 := getLeaf(cur1, stack1)
    leaf2, cur2, stack2 := getLeaf(cur2, stack2)
    for leaf1 != nil && leaf2 != nil {
        if leaf1.Val != leaf2.Val { return false }
        leaf1, cur1, stack1 = getLeaf(cur1, stack1)
        leaf2, cur2, stack2 = getLeaf(cur2, stack2)
    }
    if leaf1 == nil && leaf2 == nil { return true }
    return false
}

func getLeaf(cur *TreeNode, stack []*TreeNode) (*TreeNode, *TreeNode, []*TreeNode) {
    var ret *TreeNode
    for ret == nil {
        if cur == nil && len(stack) == 0 { break }
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack) - 1]
        if cur.Left == nil && cur.Right == nil { ret = cur }
        stack = stack[:len(stack) - 1]
        cur = cur.Right
    }
    return ret, cur, stack
}
