func numberOfArithmeticSlices(nums []int) int {
    cnt, ret := 0, 0
    for i := range nums[:len(nums)-1] {
        nums[i] = nums[i] - nums[i+1]
        if i > 0 && nums[i] == nums[i-1] {
            cnt++
        } else {
            ret += (1 + cnt) * cnt / 2
            cnt = 0
        }
    }
    ret += (1 + cnt) * cnt / 2
    return ret
}
