func search(nums []int, target int) bool {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return true
        }
    }
    return false
}

// func search(nums []int, target int) bool {
//     if len(nums) == 0 {
//         return false
//     }
//     if len(nums) == 1 {
//         return nums[0] == target
//     }
//     pivot := nums[0]
//     k := binSrchRotaIdx(nums)
//     if pivot < target {
//         return binSrch(nums[:k+1], target)
//     } else if pivot > target {
//         return binSrch(nums[k+1:len(nums)], target)
//     }
//     return true
// }

// func binSrch(nums []int, target int) bool {
//     for front, back := 0, len(nums) - 1; front <= back; {
//         mid := (front + back) / 2
//         if nums[mid] == target {
//             return true
//         } else if nums[mid] > target {
//             back = mid - 1
//         } else {
//             front = mid + 1
//         }
//     }
//     return false
// }

// func binSrchRotaIdx(nums []int) int {
//     pivot := nums[0]
//     front, back := 0, len(nums) - 1
//     for front <= back {
//         mid := (front + back) / 2
//         if nums[mid] > nums[mid + 1] {
//             return mid
//         }
//         if nums[mid] > pivot {
//             front = mid
//         } else {
//             back = mid
//         }
//     }
//     return front
// }

