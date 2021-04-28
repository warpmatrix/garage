func judgeSquareSum(c int) bool {
	lptr, rptr := 0, int(math.Sqrt(float64(c)))
	for lptr <= rptr {
		sum := lptr * lptr + rptr * rptr
		if sum == c { return true }
		if sum < c {
            lptr++ 
        } else { rptr-- }
	}
	return false
}
