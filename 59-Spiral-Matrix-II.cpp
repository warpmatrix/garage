#include <vector>
using std::vector;

vector<vector<int>> generateMatrix(int n) {
    std::vector<std::vector<int>> ret(n, std::vector<int>(n, 0));
    int index = 0;
    for (int top = 0, bottom = n - 1, left = 0, right = n - 1;
         top <= bottom && left <= right; top++, bottom--, left++, right--) {
        for (int col = left; col <= right; col++) {
            ret[top][col] = index++;
        }
        for (int row = top + 1; row <= bottom; row++) {
            ret[row][right] = index++;
        }
        for (int col = right - 1; bottom > top && col >= left; col--) {
            ret[bottom][col] = index++;
        }
        for (int row = bottom - 1; right > left && row >= top + 1; row--) {
            ret[row][left] = index++;
        }
    }
    return ret;
}