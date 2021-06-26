const r, w = 2, 3

type node struct {
    stat []int
    cost int
    estCost int
}
type priQueue []node
func (pq priQueue) Len() int {
    return len(pq)
}
func (pq priQueue) Less(lhs, rhs int) bool {
    p1, p2 := pq[lhs], pq[rhs]
    return p1.cost + p1.estCost >= p2.cost + p2.estCost
}
func (pq priQueue) Swap(lhs, rhs int) {
    pq[lhs], pq[rhs] = pq[rhs], pq[lhs]
}
func (pq *priQueue) Push(v interface{}) {
    tmp := v.(node)
    copy(tmp.stat, v.(node).stat)
    *pq = append(*pq, tmp)
}
func (pq *priQueue) Pop() interface{} {
    a := *pq
    v := a[len(a)-1]
    *pq = a[:len(a)-1]
    return v
}

func slidingPuzzle(board [][]int) int {
    tarStat := []int{1, 2, 3, 4, 5, 0}
    tarLoc := getElemLoc(tarStat)
    stat := []int{}
    for _, rowStat := range board {
        stat = append(stat, rowStat...)
    }

    h := priQueue{}
    heap.Push(&h, node{
        stat : stat, cost : 0,
        estCost : calcManhattanDist(stat, tarLoc),
    })
    visitedTbl := map[[]int]struct{}{
        stat : struct{}{},
    }
    fmt.Println(len(visitedTbl), heap)
    // for len(h) > 0 {
    //     n := heap.Pop(&h).(node)
    //     for _, neibor := range getNeibors(n.stat) {
    //         if _, visied := visitedTbl[neibor]; !visied {
    //             if reflect.DeepEqual(neibor, tarStat) {
    //                 return n.cost + 1
    //             }
    //             visitedTbl[neibor] = struct{}{}
    //             heap.Push(&h, node{
    //                 stat : neibor,
    //                 cost : n.cost + 1,
    //                 estCost : calcManhattanDist(neibor, tarStat)
    //             })
    //         }
    //     }
    // }
    return -1
}

func getNeibors(stat []int) [][]int {
    
}

func getElemLoc(stat []int) []int {
    ret := make([]int, len(stat))
    for loc, elem := range stat {
        ret[elem] = loc
    }
    return ret
}

func calcManhattanDist(stat []int, elemLoc []int) int {
    ret := 0
    for loc, elem := range stat {
        rowDist := abs(loc / w - elemLoc[elem] / w)
        colDist := abs(loc % w - elemLoc[elem] % w)
        ret += rowDist + colDist
    }
    return ret
}

func abs(num int) int {
    if num < 0 { return -num }
    return num
}
