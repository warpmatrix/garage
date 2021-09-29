func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target { return 0 }
    source, target = source - 1, target - 1
    // the routes that the node in
    nodeRidxs := make(map[int][]int)
    adjRoutes := make([]map[int]struct{}, len(routes))
    for i := range adjRoutes {
        adjRoutes[i] = make(map[int]struct{})
    }
    for rIdx, route := range routes {
        for j := range route {
            routes[rIdx][j]--
            node := routes[rIdx][j]
            for _, adjRidx := range nodeRidxs[node] {
                adjRoutes[rIdx][adjRidx] = struct{}{}
                adjRoutes[adjRidx][rIdx] = struct{}{}
            }
            nodeRidxs[node] = append(nodeRidxs[node], rIdx)
        }
    }

    queue := []int{}
    dists := make([]int, len(routes))
    for i := range dists {
        dists[i] = -1
    }
    for _, rIdx := range nodeRidxs[source] {
        dists[rIdx] = 1
        queue = append(queue, rIdx)
    }
    for len(queue) > 0 {
        rIdx1 := queue[0]
        queue = queue[1:]
        for rIdx2 := range routes {
            if _, isAdj := adjRoutes[rIdx1][rIdx2]; isAdj && dists[rIdx2] == -1 {
                dists[rIdx2] = dists[rIdx1] + 1
                queue = append(queue, rIdx2)
            }
        }
    }

    ret := math.MaxInt32
    for _, rIdx := range nodeRidxs[target] {
        if dists[rIdx] != -1 {
            ret = min(ret, dists[rIdx])
        }
    }
    if ret == math.MaxInt32 {
        ret = -1
    }
    return ret
}

func min(lhs, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}

// func numBusesToDestination(routes [][]int, source int, target int) int {
//     if source == target { return 0 }
//     source, target = source - 1, target - 1
//     // the routes that the node in
//     maxNode := -1
//     nodeRidxs := make(map[int][]int)
//     for i, route := range routes {
//         for j := range route {
//             routes[i][j]--
//             node := routes[i][j]
//             maxNode = max(maxNode, node)
//             nodeRidxs[node] = append(nodeRidxs[node], i)
//         }
//     }
//     if target > maxNode {
//         return -1
//     }
//     dists := make(map[int]int)
//     dists[source] = 0
//     queue := []int{source}
//     for len(queue) > 0 {
//         node := queue[0]
//         queue = queue[1:]
//         for _, rIdx := range nodeRidxs[node] {
//             for _, adjNode:= range routes[rIdx] {
//                 if _, calced := dists[adjNode]; calced { continue }
//                 if adjNode == target { return dists[node] + 1 }
//                 dists[adjNode] = dists[node] + 1
//                 queue = append(queue, adjNode)
//             }
//         }
//     }
//     return -1
// }

// func max(lhs, rhs int) int {
//     if rhs > lhs { return rhs }
//     return lhs
// }
