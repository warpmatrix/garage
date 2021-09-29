func findMin(nums []int) int {
	k := binSrchRotaIdx(nums)
	return nums[(k + 1) % len(nums)]
}

func binSrchRotaIdx(nums []int) int {
    pivot := nums[0]
    front, back := 0, len(nums) - 1
    for front <= back {
        mid := (front + back) / 2
        if mid == len(nums) - 1 || nums[mid] > nums[mid + 1] {
            return mid
        }
        if nums[mid] >= pivot {
            front = mid + 1
        } else {
            back = mid - 1
        }
    }
    return front
}
