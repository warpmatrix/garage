/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    stack := []*TreeNode{}
    node := root
    var newRoot, prev *TreeNode
    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack) - 1]
        if newRoot == nil {newRoot = node}
        if prev != nil {
            prev.Right = node
        }
        prev = node
        node.Left = nil
        stack = stack[:len(stack) - 1]
        node = node.Right
    }
    return newRoot
}