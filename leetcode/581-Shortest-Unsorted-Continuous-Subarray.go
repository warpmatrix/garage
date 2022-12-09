func findUnsortedSubarray(nums []int) int {
    if sort.IntsAreSorted(nums) { return 0 }
    min, max := math.MaxInt32, math.MinInt32
    left, right := -1, -1
    for i := range nums {
        if nums[i] >= max {
            max = nums[i]
        } else {
            right = i
        }
        if nums[len(nums) - i - 1] <= min {
            min = nums[len(nums) - i - 1]
        } else {
            left = len(nums) - i - 1
        }
    }
    return right - left + 1
}
