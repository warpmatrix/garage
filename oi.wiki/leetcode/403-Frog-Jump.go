func canCross(stones []int) bool {
	if stones[1] > 1 {
		return false
	}
	dist := make([]int, len(stones) - 1)
	for i:=0; i<len(stones)-1; i++ {
		dist[i] = stones[i+1] - stones[i]
	}
	nextDist := dist[len(dist) - 1]
	sum := 0
	for i := len(dist) - 2; i >= 0; i-- {
		if i == 0 && nextDist > 2 {
		    return false
		}
		if dist[i] + sum >= nextDist - 1 && dist[i] + sum <= nextDist + 1 {
			nextDist = dist[i] + sum
			sum = 0
		} else if dist[i] + sum + dist[i - 1] >= nextDist {
			nextDist += dist[i]
		} else {
			sum += dist[i]
		}
	}
	return sum == 0
}