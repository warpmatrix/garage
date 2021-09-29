func hIndex(nums []int) int {
    return len(nums) - sort.Search(len(nums), func(i int) bool {
        return nums[i] >= len(nums) - i
    })
}
