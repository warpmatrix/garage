func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n + 1)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	for _, str := range strs {
		oCnt := strings.Count(str, "1")
		zCnt := len(str) - oCnt
        
		for i := m - zCnt; i >= 0; i-- {
			for j := n - oCnt; j >= 0; j-- {
				if dp[i][j] != -1 && dp[i][j] + 1 > dp[i + zCnt][j + oCnt] {
					dp[i + zCnt][j + oCnt] = dp[i][j] + 1
				}
			}
		}
	}

	max := 0
	for _, a := range dp {
		for _, v := range a {
			if v > max { max = v }
		}
	}
	return max
}
