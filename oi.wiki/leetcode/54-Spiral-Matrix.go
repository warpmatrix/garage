func spiralOrder(matrix [][]int) []int {
    ret := make([]int, len(matrix) * len(matrix[0]))
    index := 0
    for top, left, bottom, right := 0, 0, len(matrix) - 1, len(matrix[0]) - 1; 
		left <= right && top <= bottom; 
		top, left, bottom, right = top + 1, left + 1, bottom - 1, right - 1 {
        for col := left; col <= right; col++ {
			ret[index] = matrix[top][col]
            index++
		}
		for row := top + 1; row <= bottom; row++ {
			ret[index] = matrix[row][right]
			index++
		}
		for col := right - 1; bottom > top && col >= left; col-- {
			ret[index] = matrix[bottom][col]
			index++
		}
		for row := bottom - 1; right > left && row >= top + 1; row-- {
			ret[index] = matrix[row][left]
			index++
		}
    }
    return ret
}