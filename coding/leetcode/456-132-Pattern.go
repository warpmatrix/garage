func find132pattern(nums []int) bool {
    elem2 := math.MinInt32
	elem3 := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < elem2 {
			return true
		}
		for len(elem3) > 0 && nums[i] > elem3[len(elem3) - 1] {
			elem2 = elem3[len(elem3) - 1]
			elem3 = elem3[:len(elem3) - 1]
		}
		if nums[i] > elem2 {
			elem3 = append(elem3, nums[i])
		}
	}
    return false
}
