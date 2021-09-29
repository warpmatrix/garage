func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        return strconv.Itoa(x) + strconv.Itoa(y) > strconv.Itoa(y) + strconv.Itoa(x)
    })
    if nums[0] == 0 {
        return "0"
    }
    ans := []byte{}
    for _, x := range nums {
        ans = append(ans, strconv.Itoa(x)...)
    }
    return string(ans)
}
