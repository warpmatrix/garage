func searchMatrix(matrix [][]int, target int) bool {
    row, col := len(matrix), len(matrix[0])
    head, end := 0, row * col
    for head < end {
		mid := (head + end) / 2
		if num := matrix[mid / col][mid % col]; num == target {
			return true
		} else if num < target {
			head = mid + 1
		} else {
			end = mid
		}
    }
	return false
}