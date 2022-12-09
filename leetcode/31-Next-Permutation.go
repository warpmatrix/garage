func nextPermutation(nums []int)  {
    idx1 := len(nums) - 2
    for idx1 >= 0 && nums[idx1 + 1] <= nums[idx1] {
        idx1--
    }
    if idx1 == -1 {
        reverse(nums)
        return
    }
    idx2 := len(nums) - 1
    for nums[idx2] <= nums[idx1] {
        idx2--
    }
    nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
    reverse(nums[idx1 + 1:])
}

func reverse(nums []int) []int {
    for i := 0; i * 2 < len(nums); i++ {
        tmp := len(nums) - i - 1
        nums[i], nums[tmp] = nums[tmp], nums[i]
    }
    return nums
}
