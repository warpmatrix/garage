type T struct {
    num int
}

type priQueue []T
func (pq priQueue) Len() int { return len(pq) }
func (pq priQueue) Less(lhs, rhs int) bool {
    p1, p2 := pq[lhs], pq[rhs]
    return p1.num >= p2.num
}
func (pq priQueue) Swap(lhs, rhs int) { pq[lhs], pq[rhs] = pq[rhs], pq[lhs] }
func (pq *priQueue) Push(v interface{}) { *pq = append(*pq, v.(T)) }
func (pq *priQueue) Pop() interface{} {
    a := *pq
    v := a[len(a)-1]
    *pq = a[:len(a)-1]
    return v
}
