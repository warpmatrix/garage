func checkSubarraySum(nums []int, k int) bool {
    if len(nums) < 2 { return false }
    rmdIdxs := map[int]struct{}{
        0 : struct{}{},
    }
    pRmd, rmd := nums[0] % k, nums[0] % k
    for _, num := range nums[1:] {
        rmd = (rmd + num) % k
        if _, has := rmdIdxs[rmd]; has {
            return true
        }
        rmdIdxs[pRmd], pRmd = struct{}{}, rmd
    }
    return false
}
