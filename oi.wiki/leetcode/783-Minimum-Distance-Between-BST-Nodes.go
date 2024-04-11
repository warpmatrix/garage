/**
* Definition for a binary tree node.
* type TreeNode struct {
	*     Val int
	*     Left *TreeNode
	*     Right *TreeNode
	* }
*/
	
func min(lhs int, rhs int) int {
	if lhs <= rhs { return lhs }
	return rhs
}

func minDiffInBST(root *TreeNode) int {
	ret, prev := math.MaxInt32, -1
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack) - 1]
		if prev != -1 { ret = min(ret, node.Val - prev) }
		prev = node.Val
		stack = stack[:len(stack) - 1]
		node = node.Right
	}
	return ret
}
	
// func minDiffInBST(root *TreeNode) int {
//     var dfs func(node *TreeNode)
//     prev, ret := -100000, math.MaxInt32
//     dfs = func(node *TreeNode) {
//         if node == nil {return}
//         dfs(node.Left)
//         ret = min(ret, node.Val - prev)
//         prev = node.Val
//         dfs(node.Right)
//     }
//     dfs(root)
//     return ret
// }

// func mindiffinbst(root *treenode) int {
//     var dfs func(node *treenode)
//     prev, ret := -1, math.maxint32
//     dfs = func(node *treenode) {
//         if node == nil {return}
//         dfs(node.left)
//         if prev != -1 { ret = min(ret, node.val - prev) }
//         prev = node.val
//         dfs(node.right)
//     }
//     dfs(root)
//     return ret
// }
