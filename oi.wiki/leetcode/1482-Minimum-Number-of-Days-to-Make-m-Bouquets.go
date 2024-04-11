func minDays(bloomDays []int, m int, k int) int {
    if len(bloomDays) < m * k { return -1 }
    tmp := make([]int, len(bloomDays))
    copy(tmp, bloomDays)
    sort.Ints(tmp)
    idx := sort.Search(len(bloomDays), func (idx int) bool {
        adjNum, makeNum := 0, 0
        for _, bloomDay := range bloomDays {
            if tmp[idx] >= bloomDay {
                adjNum++
                if adjNum >= k {
                    makeNum++
                    adjNum = 0
                }
            } else {
                adjNum = 0
            }
        }
        return makeNum >= m
    })
    return tmp[idx]
}

// func minDays(bloomDays []int, m int, k int) int {
//     if len(bloomDays) < m * k { return -1 }
//     tmp := make([]int, len(bloomDays))
//     copy(tmp, bloomDays)
//     sort.Ints(tmp)
//     canMake := func (day int) bool {
//         adjNum, makeNum := 0, 0
//         for _, bloomDay := range bloomDays {
//             if day >= bloomDay {
//                 adjNum++
//                 if adjNum >= k {
//                     makeNum++
//                     adjNum = 0
//                 }
//             } else {
//                 adjNum = 0
//             }
//         }
//         return makeNum >= m
//     }
//     idx := binSrch(tmp, canMake)
//     return tmp[idx]
// }

// func binSrch(arr []int, satfyFunc func (num int) bool) int {
//     min, max := 0, len(arr) - 1
//     for min < max {
//         mid := (min + max) / 2
//         if !satfyFunc(arr[mid]) {
//             min = mid + 1
//         } else {
//             max = mid
//         }
//     }
//     return min
// }

