func circularArrayLoop(nums []int) bool {
    next := func(cur, step int) int {
        return ((cur + step) % len(nums) + len(nums)) % len(nums)
    }
    for start := 0; start < len(nums); start++ {
        if nums[start] == math.MinInt32 { continue }
        ptr, step, isPos := start, 0, nums[start] > 0
        tmp := []int{}
        for {
            step = nums[ptr]
            nums[ptr] = 0
            tmp = append(tmp, ptr)
            if n := next(ptr, step); ptr == n {
                for _, idx := range tmp { nums[idx] = math.MinInt32 }
                break
            } else {
                ptr = n
            }
            if nums[ptr] == 0 { return true }
            if nums[ptr] == math.MinInt32 || isPos != (nums[ptr] > 0) {
                for _, idx := range tmp { nums[idx] = math.MinInt32 }
                break
            }
        }
    }
    return false
}
