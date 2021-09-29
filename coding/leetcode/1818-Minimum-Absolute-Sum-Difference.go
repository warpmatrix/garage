func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
    const mod = 1e9 + 7
    sorted1 := append([]int{}, nums1...)
    sort.Ints(sorted1)
    sum, maxRdc := 0, 0
    for i, v2 := range nums2 {
        diff := abs(nums1[i] - v2)
        sum += diff
        idx := sort.Search(len(sorted1), func(idx int) bool {
            return sorted1[idx] >= v2
        })
        if idx < len(sorted1) {
            maxRdc = max(maxRdc, diff - (sorted1[idx] - v2))
        }
        if idx > 0 {
            maxRdc = max(maxRdc, diff - (v2 - sorted1[idx - 1]))
        }
    }
    return (sum - maxRdc) % mod
}

func abs(num int) int {
    if num < 0 { num = -num }
    return num
}

func max(lhs, rhs int) int {
    if lhs < rhs { return rhs }
    return lhs
}
