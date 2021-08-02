/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    type location struct {
        row, col, val int
    }
    nodes := []location{}
    var dfs func(*TreeNode, int, int)
    dfs = func(node *TreeNode, row, col int) {
        if node == nil { return }
        nodes = append(nodes, location{row, col, node.Val})
        dfs(node.Left, row + 1, col - 1)
        dfs(node.Right, row + 1, col + 1)
    }
    dfs(root, 0, 0)
    sort.Slice(nodes, func (i, j int) bool {
        return nodes[i].col < nodes[j].col || nodes[i].col == nodes[j].col && (
            nodes[i].row < nodes[j].row || nodes[i].row == nodes[j].row &&
                nodes[i].val < nodes[j].val)
    })
    ret, prev := [][]int{}, math.MinInt32
    for _, node := range nodes {
        if node.col != prev {
            ret = append(ret, []int{})
            prev = node.col
        }
        idx := len(ret) - 1
        ret[idx] = append(ret[idx], node.val)
    }
    return ret
}
