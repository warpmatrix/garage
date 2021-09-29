func majorityElement(nums []int) int {
    if len(nums) == 0 { return -1 }
    majElem := nums[0]
    cnt := 1
    for _, num := range nums[1:] {
        if num == majElem {
            cnt++
        } else {
            cnt--
            if cnt == -1 {
                majElem = num
                cnt = 1
            }
        }
    }
    cnt = 0
    for _, num := range nums {
        if num == majElem {
            cnt++
        }
    }
    if cnt * 2 <= len(nums) {
        return -1
    }
    return majElem
}
