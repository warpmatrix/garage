func maxSubArray(nums []int) int {
    preSum, min, ret := 0, 0, math.MinInt32
    for _, num := range nums {
        preSum += num
        ret = max(ret, preSum - min)
        if preSum < min {
            min = preSum
        }
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}

// 使用数组的空间维护 min 和 preSum 变量，虽然易于理解，不过可读性好像变差了
// func maxSubArray(nums []int) int {
//     max:=nums[0]
//     for i:=1;i<len(nums);i++{
//         if nums[i]+nums[i-1]>nums[i]{
//             nums[i] = nums[i]+nums[i-1]
//         }
//         if nums[i]>max{
//             max=nums[i]
//         } 
//     }
//     return max
// }

