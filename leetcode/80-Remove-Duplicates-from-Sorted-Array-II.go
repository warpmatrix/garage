func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    idx, cnt := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            cnt++
        } else {
            cnt = 1
        }
        if cnt <= 2 {
            nums[idx] = nums[i]
            idx++
        }
    }
    return idx
}