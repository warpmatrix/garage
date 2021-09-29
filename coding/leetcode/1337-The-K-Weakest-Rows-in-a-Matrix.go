func kWeakestRows(mat [][]int, k int) []int {
    type pair struct {
        num, idx int
    }
    slice := make([]pair, len(mat))
    for rIdx, row := range mat {
        cnt := 0
        for _, elem := range row {
            if elem == 0 { break }
            cnt++
        }
        slice[rIdx] = pair{num:cnt, idx:rIdx}
    }
    sort.Slice(slice, func (i, j int) bool {
        return slice[i].num < slice[j].num ||
            slice[i].num == slice[j].num && slice[i].idx < slice[j].idx
    })
    ret := make([]int, k)
    for i := range ret {
        ret[i] = slice[i].idx
    }
    return ret
}
