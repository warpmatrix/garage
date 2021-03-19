func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	idx := 1
	for top, left, bottom, right := 0, 0, n - 1, n - 1; 
		left <= right && top <= bottom; 
		top, left, bottom, right = top + 1, left + 1, bottom - 1, right - 1 {
        for col := left; col <= right; col++ {
			ret[top][col] = idx
            idx++
		}
		for row := top + 1; row <= bottom; row++ {
			ret[row][right] = idx
			idx++
		}
		for col := right - 1; bottom > top && col >= left; col-- {
			ret[bottom][col] = idx
			idx++
		}
		for row := bottom - 1; right > left && row >= top + 1; row-- {
			ret[row][left] = idx
			idx++
		}
    }
	return ret
}