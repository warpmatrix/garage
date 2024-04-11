func numSquares(n int) int {
    if isSquare(n) { return 1 }
    if satAns4(n) { return 4 }
    if satAns2(n) { return 2 }
    return 3
}

func satAns4(n int) bool {
    for n % 4 == 0 {
        n /= 4
    }
    // n % 8 == 7
    return n & 7 == 7
}

func satAns2(n int) bool {
    for i := 1; i * i <= n / 2; i++ {
        if isSquare(n - i * i) {
            return true
        }
    }
    return false
}

func isSquare(n int) bool {
    root := int(math.Sqrt(float64(n)))
    return n == root* root
}

// func numSquares(n int) int {
//     squares := make([]int, int(math.Sqrt(float64(n))))
//     for i := range squares {
//         squares[i] = (i + 1) * (i + 1)
//     }
//     dp := make([]int, n + 1)
//     for i := 1; i <= n; i++ {
//         dp[i] = math.MaxInt32
//         for _, square := range squares {
//             if square > i { break }
//             dp[i] = min(dp[i], dp[i - square] + 1)
//         }
//     }
//     return dp[n]
// }

// func min (lhs, rhs int) int {
//     if rhs < lhs { return rhs }
//     return lhs
// }
