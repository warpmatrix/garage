func search(nums []int, target int) int {
    head := sort.SearchInts(nums, target)
    tail := sort.SearchInts(nums, target + 1)
    return tail - head
}
