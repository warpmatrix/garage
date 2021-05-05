func deleteAndEarn(nums []int) int {
	if len(nums) == 0 { return 0 }
    sort.Ints(nums)
    idxs := []int{1}
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
			idxs[len(idxs) - 1]++
		} else {
            idxs = append(idxs, idxs[len(idxs) - 1] + 1)
        }
    }
	dp := make([]int, len(idxs) + 1)
	dp[0], dp[1] = 0, nums[idxs[0] - 1] * idxs[0]
	for i := 2; i <= len(idxs); i++ {
		if nums[idxs[i] - 1] != nums[idxs[i - 1] - 1] + 1 {
			dp[i] = dp[i - 1] + nums[idxs[i] - 1] * (idxs[i] - idxs[i - 1])
		} else {
			dp[i] = max(dp[i - 1], dp[i - 2] + nums[idxs[i] - 1] * (idxs[i] - idxs[i - 1]))
		}
	}
    return dp[len(idxs)]
}