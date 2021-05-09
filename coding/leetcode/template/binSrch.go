// return the min index of array that satisfy the func
func binSrch(arr []int, satfyFunc func (num int) bool) int {
    min, max := 0, len(arr) - 1
    for min < max {
        mid := (min + max) / 2
        if !satfyFunc(arr[mid]) {
            min = mid + 1
        } else {
            max = mid
        }
    }
    return min
}