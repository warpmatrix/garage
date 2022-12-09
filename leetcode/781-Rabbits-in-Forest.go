func numRabbits(answers []int) int {
    cnts := map[int]int{}
	for i := 0; i < len(answers); i++ {
		cnts[answers[i]]++
	}
	ret := 0
	for idx, cnt := range cnts {
		ans += (cnt / (idx + 1) + 1) * cnt;
	}
	return ret
}