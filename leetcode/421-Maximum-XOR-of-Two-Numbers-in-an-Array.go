func findMaximumXOR(a []int) int {
    ret, mask := 0, 0
    for bit := 30; bit >= 0; bit-- {
		mask |= 1 << bit
		try := ret | 1 << bit
		leftPart := map[int]struct{} { a[0] & mask: {} }
		for _, v := range a[1:] {
			if _, has := leftPart[(v & mask) ^ try]; has {
				ret = try
				break
			}
			leftPart[v & mask] = struct{}{}
		}
	}
	return ret
}
