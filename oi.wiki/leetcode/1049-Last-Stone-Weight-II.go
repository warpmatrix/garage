func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, weight := range stones {
		sum += weight
	}
	dp := make([]bool, sum / 2 + 1)
	dp[0] = true
	for _, weight := range stones {
		for sumWei := len(dp) - 1; sumWei >= 0; sumWei-- {
			if dp[sumWei] == true && sumWei + weight < len(dp) {
				dp[sumWei + weight] = true
			}
		}
	}
	negMax := 0
	for sumWei := len(dp) - 1; sumWei >= 0; sumWei-- {
		if dp[sumWei] == true {
            negMax = sumWei
			break
		}
	}
	return sum - 2 * negMax
}
