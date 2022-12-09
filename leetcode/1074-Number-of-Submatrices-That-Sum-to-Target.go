func numSubmatrixSumTarget(matrix [][]int, target int) int {
    ret := 0
    for ub := range matrix {
        sum := make([]int, len(matrix[0]))
        for _, row := range matrix[ub:] {
            for col, v := range row {
                sum[col] += v
            }
            ret += subarraySum(sum, target)
        }
    }
    return ret
}

func subarraySum(nums []int, k int) int {
    cnts := map[int]int{0:1}
    preSum, ret := 0, 0
    for _, num := range nums {
        preSum += num
        ret += cnts[preSum - k]
        cnts[preSum]++
    }
    return ret
}
