func shortestPathLength(graph [][]int) int {
    type pair struct { mask, node int }
    type tuple struct {
        p pair
        dist int
    }
    q, visited := []tuple{}, map[pair]struct{}{}
    for i := range graph {
        q = append(q, tuple{pair{1 << i, i}, 0})
        visited[pair{1 << i, i}] = struct{}{}
    }
    for len(q) > 0 {
        front := q[0]
        if front.p.mask == 1 << len(graph) - 1 { return front.dist }
        q = q[1:]
        for _, adjNode := range graph[front.p.node] {
            newSt := pair{front.p.mask | 1 << adjNode, adjNode}
            if _, isVisited := visited[newSt]; isVisited { continue }
            q = append(q, tuple{newSt, front.dist + 1})
            visited[newSt] = struct{}{}
        }
    }
    return 0
}
