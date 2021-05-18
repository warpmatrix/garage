// func countTriplets(arr []int) int {
//     if len(arr) == 0 { return 0 }
//     var preXor, ret int
//     // record count and sum of index with the same preXor
//     cnts, sumIdxs := map[int]int{0:1}, map[int]int{0:0}
//     for i := range arr {
//         preXor ^= arr[i]
//         if cnt, ok := cnts[preXor]; ok {
//             ret += i * cnt - sumIdxs[preXor]
//         }
//         cnts[preXor]++
//         sumIdxs[preXor] += i
//     }
//     return ret
// }

func countTriplets(arr []int) int {
    if arr == nil || len(arr) == 0 { return 0 }
    preXors := make([]int, len(arr) + 1)
    for i := range arr {
        preXors[i + 1] = preXors[i] ^ arr[i]
    }
    ret := 0
    for i := range preXors[:len(preXors) - 1] {
        for j := i + 1; j < len(preXors) - 1; j++  {
            if preXors[i] == preXors[j + 1] {
                ret += j - i
            }
        }
    }
    return ret
}
