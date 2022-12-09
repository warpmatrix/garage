func peakIndexInMountainArray(arr []int) int {
    return sort.Search(len(arr) - 1, func(idx int) bool {
        return arr[idx] > arr[idx + 1]
    })
}
