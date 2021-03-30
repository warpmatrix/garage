func findMaxLength(nums []int) int {
	min, max := make(map[int]int), make(map[int]int)
    min[0] = -1
	cnt, ret := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt--
		}else {
			cnt++
		}
		if _, ok := min[cnt]; !ok {
			min[cnt] = i
		} else {
			max[cnt] = i
			if max[cnt] - min[cnt] > ret {
				ret = max[cnt] - min[cnt]
			}
		}
	}
	return ret
}