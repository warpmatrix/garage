/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    min := root.Val
    ret := -1
    var f func(*TreeNode)
    f = func(node *TreeNode) {
        if node == nil || ret != -1 && node.Val >= ret {
            return
        }
        if node.Val > min {
            ret = node.Val
        }
        f(node.Left)
        f(node.Right)
    }
    f(root)
    return ret
}
