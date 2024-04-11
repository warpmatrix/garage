func leastBricks(wall [][]int) int {
    edgeNum := make(map[int]int, 1000)
    for _, widths := range wall {
        sum := 0
        for i := 0; i < len(widths) - 1; i++ {
            sum += widths[i]
            edgeNum[sum]++
        }
    }
    max := 0
    for _, num := range edgeNum {
        if num > max { max = num }
    }
    return len(wall) - max
}
