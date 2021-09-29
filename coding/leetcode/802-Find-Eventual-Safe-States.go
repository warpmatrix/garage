func eventualSafeNodes(graph [][]int) []int {
    revGrpha := make([][]int, len(graph))
    outDeg := make([]int, len(graph))
    queue := []int{}
    for from, tos := range graph {
        outDeg[from] = len(graph[from])
        if outDeg[from] == 0 {
            queue = append(queue, from)
            continue
        }
        for _, to := range tos {
            revGrpha[to] = append(revGrpha[to], from)
        }
    }

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for _, inNode := range revGrpha[node] {
            outDeg[inNode]--
            if outDeg[inNode] == 0 {
                queue = append(queue, inNode)
            }
        }
    }
    ret := []int{}
    for node, outDeg := range outDeg {
        if outDeg == 0 {
            ret = append(ret, node)
        } 
    }
    return ret
}
