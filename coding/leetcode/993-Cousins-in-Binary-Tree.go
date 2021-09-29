/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    depths, parents, founds := make([]int, 2), make([]*TreeNode, 2), make([]bool, 2)
    targets := []int{x, y}
    var dfs func(node *TreeNode, depth int, parent *TreeNode)
    dfs = func(node *TreeNode, depth int, parent *TreeNode) {
        if founds[0] && founds[1] { return }
        if node == nil { return }
        for i := range targets {
            if node.Val == targets[i] {
                founds[i] = true
                depths[i] = depth
                parents[i] = parent
            }
        }
        dfs(node.Left, depth + 1, node)
        dfs(node.Right, depth + 1, node)
    }
    dfs(root, 0, nil)
    return depths[0] == depths[1] && parents[0] != parents[1]
}

// func isCousins(root *TreeNode, x int, y int) bool {
//     if root == nil || root.Val == x || root.Val == y {
//         return false
//     }
//     queue, prev := []*TreeNode{root}, []*TreeNode{}
//     // depth iteration
//     for len(queue) > 0 {
//         prev, queue = queue, []*TreeNode{}
//         found := false
//         for _, node := range prev {
// 			if node == nil { continue }
//             check := func(node *TreeNode, target int) bool {
//                 return node != nil && node.Val == target
//             }
//             foundX := check(node.Left, x) || check(node.Right, x)
//             foundY := check(node.Left, y) || check(node.Right, y)
//             if foundX && foundY { return false }
//             if found && (foundX || foundY) { return true }
//             found = found || foundX || foundY
// 			queue = append(queue, node.Left)
// 			queue = append(queue, node.Right)
//         }
//         if found { return false }
//     }
//     return false
// }