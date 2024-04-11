func singleNumber(nums []int) int {
	st1, st0 := 0, 0
	for _, num := range nums {
		st1, st0 = (^st1 & st0 & num) | (st1 & ^st0 & ^num),
				   (^st1 & ^st0 & num) | (^st1 & st0 & ^num)
	}
	return st0
}
