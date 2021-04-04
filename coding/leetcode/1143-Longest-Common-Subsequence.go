func max(lhs int, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func longestCommonSubsequence(text1 string, text2 string) int {
	lcs := make([][]int, len(text1) + 1)
	for i := 0; i <= len(text1); i++ {
		lcs[i] = make([]int, len(text2) + 1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i - 1] == text2[j - 1] {
				lcs[i][j] = lcs[i - 1][j - 1] + 1
			} else {
				lcs[i][j] = max(lcs[i - 1][j], lcs[i][j - 1])
			}
		}
	}
	return lcs[len(text1)][len(text2)]
}
