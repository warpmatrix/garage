func subsetsWithDup(nums []int) [][]int {
	cnts := make(map[int]int)
	keys := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := cnts[nums[i]]; !ok {
			keys = append(keys, nums[i])
		}
		cnts[nums[i]]++
	}
	ret := [][]int{}
	cur := []int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(keys) {
			ret = append(ret, append([]int(nil), cur...))
            return
		}
		for i := 0; i <= cnts[keys[idx]]; i++ {
            for j := 0; j < i; j++ {
			    cur = append(cur, keys[idx])
            }
			dfs(idx + 1)
            cur = cur[:len(cur) - i]
		}
	}
	dfs(0)
	return ret
}