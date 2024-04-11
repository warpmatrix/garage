func snakesAndLadders(board [][]int) int {
    step := make([]int, len(board) * len(board[0]))
    for i := range step[1:] {
        step[1 + i] = -1
    }
    queue := []int{0}
    for len(queue) > 0 {
        idx := queue[0]
        queue = queue[1:]
        for detIdx := 1; detIdx <= 6; detIdx++ {
            dest := idx + detIdx
            if !(dest < len(step)) { break }
            r, c := idx2rw(dest, len(board), len(board[0]))
            if board[r][c] != -1 {
                dest = board[r][c] - 1
            }
            if step[dest] != -1 { continue }
            step[dest] = step[idx] + 1
            queue = append(queue, dest)
        }
    }
    return step[len(step) - 1]
}

func idx2rw(idx, row, col int) (int, int) {
    rIdx := idx / col
    cIdx := idx % col
    if rIdx % 2 == 1 {
        cIdx = col - cIdx - 1
    }
    rIdx = row - rIdx - 1
    return rIdx, cIdx
}
