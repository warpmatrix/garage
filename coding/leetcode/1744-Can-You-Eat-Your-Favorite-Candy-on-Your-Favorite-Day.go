func canEat(candiesCount []int, queries [][]int) []bool {
    // type, day, dCap
    preCnts := make([]int, len(candiesCount))
    preCnts[0] = candiesCount[0]
    for i, cnt := range candiesCount[1:] {
        preCnts[i + 1] = preCnts[i] + cnt
    }
    ret := make([]bool, len(queries))
    for i, q := range queries {
        favType, day, cap := q[0], q[1], q[2]
        minEat, maxEat := day + 1, cap * (day + 1)
        minCnt, maxCnt := 1, preCnts[favType]
        if favType > 0 { minCnt = preCnts[favType - 1] + 1 }
        ret[i] = !(minEat > maxCnt || maxEat < minCnt)
    }
    return ret
}
