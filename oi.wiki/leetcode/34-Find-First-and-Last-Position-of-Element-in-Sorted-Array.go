func searchRange(nums []int, target int) []int {
    head := sort.SearchInts(nums, target)
    if head == len(nums) || nums[head] != target {
        return []int{-1, -1}
    }
    tail := sort.SearchInts(nums, target + 1) - 1
    return []int{head, tail}
}
