func removeDuplicates(nums []int) int {
    len := 0
    for _, v := range nums {
        if len > 0 && nums[len - 1] == v { continue }
        nums[len] = v
        len++
    }
    return len
}