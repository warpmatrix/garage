func findMaxLength(nums []int) int {
	minIdxs := make(map[int]int)
    minIdxs[0] = -1
	cnt, ret := 0, 0
	for i, num := range nums {
		cnt += (num << 1) - 1
		if minIdx, exist := minIdxs[cnt]; !exist {
			minIdxs[cnt] = i
		} else {
            ret = max(ret, i - minIdx)
		}
	}
	return ret
}

func max(lhs, rhs int) int {
	if lhs < rhs { return rhs }
	return lhs
}