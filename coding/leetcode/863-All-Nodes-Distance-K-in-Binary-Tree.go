/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    ret := []int{}
    var findDesc func(*TreeNode, int, int)
    findDesc = func(node *TreeNode, depth, trgDep int) {
        if node == nil { return }
        if depth == trgDep {
            ret = append(ret, node.Val)
            return
        }
        if node.Left != target { findDesc(node.Left, depth + 1, trgDep) }
        if node.Right != target { findDesc(node.Right, depth + 1, trgDep) }
    }
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return -1 }
        if node == target {
            findDesc(node, 0, k)
            return 0
        }
        if dist := dfs(node.Left) + 1; dist > 0 {
            if k - dist == 0 {
                ret = append(ret, node.Val)
            } else {
                findDesc(node.Right, 1, k - dist)
            }
            return dist
        } else if dist := dfs(node.Right) + 1; dist > 0 {
            if k - dist == 0 {
                ret = append(ret, node.Val)
            } else {
                findDesc(node.Left, 1, k - dist)
            }
            return dist
        }
        return -1
    }
    dfs(root)
    return ret
}
