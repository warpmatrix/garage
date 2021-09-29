func isCovered(ranges [][]int, left int, right int) bool {
    left, right = left - 1, right - 1
    diffs := make([]int, 51)
    for _, r := range ranges {
        r[0], r[1] = r[0] - 1, r[1] - 1
        diffs[r[0]]++
        diffs[r[1] + 1]--
    }
    cnt := 0
    for i := range diffs[:right + 1] {
        cnt += diffs[i]
        if left <= i && cnt <= 0 {
            return false
        }
    }
    return true
}
