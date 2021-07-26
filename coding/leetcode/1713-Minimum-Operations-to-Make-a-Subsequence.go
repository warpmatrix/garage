func minOperations(target []int, arr []int) int {
    idxs := map[int][]int{}
    for _, num := range target {
        idxs[num] = []int{}
    }
    for idx, num := range arr {
        if _, isTgt := idxs[num]; isTgt {
            idxs[num] = append(idxs[num], idx)
        }
    }
    minElem := []int{}
    for _, num := range target {
        for i := len(idxs[num]) - 1; i >= 0; i-- {
            idx := sort.SearchInts(minElem, idxs[num][i])
            if idx == len(minElem) {
                minElem = append(minElem, idxs[num][i])
            } else {
                minElem[idx] = idxs[num][i]
            }
        }
    }
    return len(target) - len(minElem)
}
