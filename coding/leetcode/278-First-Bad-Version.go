/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    return 1 + sort.Search(n, func(versionm1 int) bool {
        return isBadVersion(versionm1 + 1)
    })
}

// func firstBadVersion(n int) int {
//     return binSrch(1, n, func (num int) int {
//         return num
//     }, isBadVersion)
// }

// func binSrch(min, max int, mapFunc func(num int) int,
//     satfyFunc func(num int) bool) int {
//     for min < max {
//         mid := (min + max) / 2
//         if !satfyFunc(mapFunc(mid)) {
//             min = mid + 1
//         } else {
//             max = mid
//         }
//     }
//     return min
// }
