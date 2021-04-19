func removeElement(nums []int, val int) int {
    len := 0
    for _, v := range nums {
        if val == v { continue }
        nums[len] = v
        len++
    }
    return len
}
