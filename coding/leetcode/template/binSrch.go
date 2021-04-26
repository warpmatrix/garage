// return the min index of target
func binSrch(min int, max int, func fun(num int) int, target int) int {
	for min < max {
		mid := (min + max) / 2
		if fun(mid) < target {
			min = mid + 1
		} else {
			max = mid 
		}
	}
	return min
}