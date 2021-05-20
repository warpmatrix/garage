type pair struct {
    word string
    cnt int
}

type priQueue []pair
func (pq priQueue) Len() int { return len(pq) }
func (pq priQueue) Less(i1, i2 int) bool {
    p1, p2 := pq[i1], pq[i2]
    return p1.cnt > p2.cnt || p1.cnt == p2.cnt && p1.word < p2.word
}
func (pq priQueue) Swap(i1, i2 int) { pq[i1], pq[i2] = pq[i2], pq[i1] }
func (pq *priQueue) Push(v interface{}) { *pq = append(*pq, v.(pair)) }
func (pq *priQueue) Pop() interface{} {
    a := *pq
    v := a[len(a)-1]
    *pq = a[:len(a)-1]
    return v
}

func topKFrequent(words []string, k int) []string {
    cnts := make(map[string]int)
    for _, word := range words {
        cnts[word]++
    }
    pq := &priQueue{}
    for word, cnt := range cnts {
        heap.Push(pq, pair{word, cnt})
        // pq.Push(pair{word, cnt})
    }
    ret := make([]string, k)
    for i := 0; i < k; i++ {
        ret[i] = heap.Pop(pq).(pair).word
        // ret[i] = pq.Pop().(pair).word
    }
    return ret
}

// func topKFrequent(words []string, k int) []string {
//     cnts := make(map[string]int)
//     for _, word := range words {
//         cnts[word]++
//     }
//     ret := []string{}
//     for unqWord := range cnts {
//         ret = append(ret, unqWord)
//     }
//     sort.Slice(ret, func (i1 int, i2 int) bool {
//         w1, w2 := ret[i1], ret[i2]
//         return cnts[w1] > cnts[w2] || cnts[w1] == cnts[w2] && w1 < w2
//     })
//     ret = ret[:k]
//     return ret
// }