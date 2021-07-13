type st struct {
    key int
}

type hp []st
func (h hp) Len() int {
    return len(h)
}
func (h hp) Less(idx1, idx2 int) bool {
    return h[idx1].key > h[idx2].key
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
