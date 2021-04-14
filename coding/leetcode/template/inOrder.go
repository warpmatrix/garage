func inOrder(root *TreeNode) {
	stack := []*TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack) - 1]
		fmt.Prinln(node)
        stack = stack[:len(stack) - 1]
        node = node.Right
    }
}