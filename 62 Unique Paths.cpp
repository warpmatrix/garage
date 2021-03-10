#include <memory.h>

#include <algorithm>

// int uniquePaths(int m, int n) {
//     int uniPaths[m][n];
//     for (int i = 0; i < m; i++) {
//         for (int j = 0; j < n; j++) {
//             if (i == 0 || j == 0) {
//                 uniPaths[i][j] = 1;
//             } else {
//                 uniPaths[i][j] = uniPaths[i - 1][j] + uniPaths[i][j - 1];
//             }
//         }
//     }
//     return uniPaths[m - 1][n - 1];
// }

int uniquePaths(int m, int n) {
    // return (m + n - 2)! / ((m - 1)! (n - 1)!)
    int tol = m + n - 2;
    int min = std::min(m - 1, n - 1);
    unsigned long long int ret = 1;
    for (int i = 0; i < min; i++) {
        ret *= (tol - i + 1) / i;
    }
    return ret;
}
