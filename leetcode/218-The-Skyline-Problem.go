type st struct {
    right, height int
}
type hp []st
func (h hp) Len() int {
    return len(h)
}
func (h hp) Less(idx1, idx2 int) bool {
    return h[idx1].height > h[idx2].height
}
func (h hp) Swap(idx1, idx2 int) {
    h[idx1], h[idx2] = h[idx2], h[idx1]
}
func (h *hp) Push(v interface{}) {
    *h = append(*h, v.(st))
}
func (h *hp) Pop() interface{} {
    a := *h
    v := a[len(a)-1]
    *h = a[:len(a)-1]
    return v
}

func getSkyline(buildings [][]int) [][]int {
    edges := make([]int, 2 * len(buildings))
    for i, building := range buildings {
        edges[2 * i], edges[2 * i + 1] = building[0], building[1]
    }
    sort.Ints(edges)
    idx, h, ret := 0, hp{}, [][]int{}
    for _, loc := range edges {
        // there may be more then one buildings which have the same left
        for idx < len(buildings) && buildings[idx][0] <= loc {
            heap.Push(&h, st{buildings[idx][1], buildings[idx][2]})
            idx++
        }
        for len(h) > 0 && h[0].right <= loc {
            heap.Pop(&h)
        }
        maxHeight := 0
        if len(h) > 0 { maxHeight = h[0].height }
        if len(ret) == 0 || ret[len(ret) - 1][1] != maxHeight {
            ret = append(ret, []int{loc, maxHeight})
        }
    }
    return ret
}
