func restoreArray(adjacentPairs [][]int) []int {
    adjNodes := map[int][]int{}
    for _, pair := range adjacentPairs {
        adjNodes[pair[0]] = append(adjNodes[pair[0]], pair[1])
        adjNodes[pair[1]] = append(adjNodes[pair[1]], pair[0])
    }
    var node, prev int
    for key, adjNode := range adjNodes {
        if len(adjNode) == 1 {
            node, prev = key, key
            break
        }
    }
    ret := []int{ node }
    for len(ret) < len(adjacentPairs) + 1 {
        for _, adjNode := range adjNodes[node] {
            if adjNode == prev { continue }
            prev, node = node, adjNode
            break
        }
        ret = append(ret, node)
    }
    return ret
}
